<script setup lang="ts">
import { computed, onMounted, reactive, ref } from "vue";
import {
  createStudent,
  deleteStudent,
  getStudentDetail,
  getStudentList,
  updateStudent,
  updateStudentScore,
} from "@/api/student";
import type { Student } from "@/types/student";

const loading = ref(false);
const saving = ref(false);
const students = ref<Student[]>([]);
const activeId = ref("");

const notice = ref<{ type: "success" | "error"; text: string } | null>(null);

const createForm = reactive({
  id: "",
  name: "",
  age: 18,
  gender: "male",
});

const editForm = reactive({
  id: "",
  name: "",
  age: 18,
  gender: "male",
});

const scoreForm = reactive({
  studentId: "",
  subject: "",
  score: 60,
});

const filters = reactive({
  keyword: "",
  gender: "all",
  mode: "all",
  topN: 3,
});

function averageScore(student: Student): number {
  const values = Object.values(student.scores || {});
  if (values.length === 0) {
    return 0;
  }
  const sum = values.reduce((acc, value) => acc + value, 0);
  return Number((sum / values.length).toFixed(2));
}

function hasFailedSubject(student: Student): boolean {
  return Object.values(student.scores || {}).some((value) => value < 60);
}

const normalizedStudents = computed(() => {
  return students.value.map((item) => ({
    ...item,
    scores: item.scores || {},
  }));
});

const visibleStudents = computed(() => {
  let list = [...normalizedStudents.value];

  if (filters.keyword.trim()) {
    const keyword = filters.keyword.trim().toLowerCase();
    list = list.filter(
      (item) =>
        item.id.toLowerCase().includes(keyword) ||
        item.name.toLowerCase().includes(keyword),
    );
  }

  if (filters.gender !== "all") {
    list = list.filter((item) => item.gender.toLowerCase() === filters.gender);
  }

  list.sort((a, b) => averageScore(b) - averageScore(a));

  if (filters.mode === "failed") {
    list = list.filter(hasFailedSubject);
  }

  if (filters.mode === "top") {
    const n = Math.max(1, Math.floor(Number(filters.topN) || 1));
    list = list.slice(0, n);
  }

  return list;
});

const activeStudent = computed(() => {
  return normalizedStudents.value.find((item) => item.id === activeId.value) || null;
});

function showNotice(type: "success" | "error", text: string) {
  notice.value = { type, text };
  window.setTimeout(() => {
    if (notice.value?.text === text) {
      notice.value = null;
    }
  }, 2800);
}

function fillEditForm(student: Student) {
  editForm.id = student.id;
  editForm.name = student.name;
  editForm.age = student.age;
  editForm.gender = student.gender;
  scoreForm.studentId = student.id;
}

async function loadStudents() {
  loading.value = true;
  try {
    const res = await getStudentList();
    students.value = res.data;

    if (!activeId.value && students.value.length > 0) {
      activeId.value = students.value[0]?.id || "";
    }

    if (activeId.value) {
      const exists = students.value.some((item) => item.id === activeId.value);
      if (!exists) {
        activeId.value = students.value[0]?.id || "";
      }
    }

    if (activeId.value) {
      const detail = await getStudentDetail(activeId.value);
      const index = students.value.findIndex((item) => item.id === activeId.value);
      if (index >= 0) {
        students.value[index] = detail.data;
      }
      fillEditForm(detail.data);
    }
  } catch (error: any) {
    showNotice("error", error.message || "Load students failed");
  } finally {
    loading.value = false;
  }
}

async function selectStudent(id: string) {
  activeId.value = id;
  try {
    const res = await getStudentDetail(id);
    const index = students.value.findIndex((item) => item.id === id);
    if (index >= 0) {
      students.value[index] = res.data;
    }
    fillEditForm(res.data);
  } catch (error: any) {
    showNotice("error", error.message || "Load student detail failed");
  }
}

async function handleCreate() {
  if (!createForm.id.trim() || !createForm.name.trim()) {
    showNotice("error", "ID and name are required");
    return;
  }

  saving.value = true;
  try {
    await createStudent({
      id: createForm.id.trim(),
      name: createForm.name.trim(),
      age: Number(createForm.age),
      gender: createForm.gender,
    });

    createForm.id = "";
    createForm.name = "";
    createForm.age = 18;
    createForm.gender = "male";

    await loadStudents();
    showNotice("success", "Student created");
  } catch (error: any) {
    showNotice("error", error.message || "Create failed");
  } finally {
    saving.value = false;
  }
}

async function handleUpdateStudent() {
  if (!editForm.id) {
    showNotice("error", "Please select a student first");
    return;
  }

  saving.value = true;
  try {
    await updateStudent(editForm.id, {
      name: editForm.name.trim(),
      age: Number(editForm.age),
      gender: editForm.gender,
    });

    await selectStudent(editForm.id);
    await loadStudents();
    showNotice("success", "Student updated");
  } catch (error: any) {
    showNotice("error", error.message || "Update failed");
  } finally {
    saving.value = false;
  }
}

async function handleUpdateScore() {
  if (!scoreForm.studentId || !scoreForm.subject.trim()) {
    showNotice("error", "Student and subject are required");
    return;
  }

  saving.value = true;
  try {
    await updateStudentScore(scoreForm.studentId, scoreForm.subject.trim(), {
      score: Number(scoreForm.score),
    });

    await selectStudent(scoreForm.studentId);
    await loadStudents();

    scoreForm.subject = "";
    scoreForm.score = 60;

    showNotice("success", "Score updated");
  } catch (error: any) {
    showNotice("error", error.message || "Update score failed");
  } finally {
    saving.value = false;
  }
}

async function handleDelete(id: string) {
  const ok = window.confirm(`Delete student ${id}?`);
  if (!ok) {
    return;
  }

  saving.value = true;
  try {
    await deleteStudent(id);
    if (activeId.value === id) {
      activeId.value = "";
    }
    await loadStudents();
    showNotice("success", "Student deleted");
  } catch (error: any) {
    showNotice("error", error.message || "Delete failed");
  } finally {
    saving.value = false;
  }
}

onMounted(() => {
  loadStudents();
});
</script>

<template>
  <div class="page">
    <div class="glow glow-a"></div>
    <div class="glow glow-b"></div>

    <main class="board">
      <header class="hero">
        <p class="tag">EduCore Dashboard</p>
        <h1>Student Studio</h1>
        <p class="sub">
          Create, edit, score, and rank students in one focused workspace.
        </p>
      </header>

      <p v-if="notice" class="notice" :class="notice.type">
        {{ notice.text }}
      </p>

      <section class="grid">
        <article class="card">
          <h2>Create Student</h2>
          <div class="form-grid">
            <label>
              Student ID
              <input v-model="createForm.id" placeholder="e.g. S1001" />
            </label>
            <label>
              Name
              <input v-model="createForm.name" placeholder="e.g. Ada" />
            </label>
            <label>
              Age
              <input v-model.number="createForm.age" type="number" min="1" />
            </label>
            <label>
              Gender
              <select v-model="createForm.gender">
                <option value="male">male</option>
                <option value="female">female</option>
              </select>
            </label>
          </div>
          <button class="btn primary" :disabled="saving" @click="handleCreate">
            {{ saving ? "Saving..." : "Create" }}
          </button>
        </article>

        <article class="card">
          <h2>Filters & Ranking</h2>
          <div class="form-grid compact">
            <label>
              Search
              <input v-model="filters.keyword" placeholder="ID or name" />
            </label>
            <label>
              Gender
              <select v-model="filters.gender">
                <option value="all">all</option>
                <option value="male">male</option>
                <option value="female">female</option>
              </select>
            </label>
            <label>
              Mode
              <select v-model="filters.mode">
                <option value="all">all students</option>
                <option value="failed">failed only</option>
                <option value="top">top N</option>
              </select>
            </label>
            <label>
              Top N
              <input v-model.number="filters.topN" type="number" min="1" />
            </label>
          </div>
          <button class="btn ghost" :disabled="loading" @click="loadStudents">
            {{ loading ? "Refreshing..." : "Refresh data" }}
          </button>
        </article>
      </section>

      <section class="grid lower">
        <article class="card list-card">
          <div class="head-row">
            <h2>Students</h2>
            <span class="badge">{{ visibleStudents.length }}</span>
          </div>

          <div class="table-wrap">
            <table>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Name</th>
                  <th>Age</th>
                  <th>Gender</th>
                  <th>Avg</th>
                  <th>Action</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="item in visibleStudents"
                  :key="item.id"
                  :class="{ active: item.id === activeId }"
                >
                  <td>{{ item.id }}</td>
                  <td>{{ item.name }}</td>
                  <td>{{ item.age }}</td>
                  <td>{{ item.gender }}</td>
                  <td>{{ averageScore(item).toFixed(2) }}</td>
                  <td class="actions">
                    <button class="btn mini" @click="selectStudent(item.id)">View</button>
                    <button class="btn mini danger" @click="handleDelete(item.id)">
                      Delete
                    </button>
                  </td>
                </tr>
                <tr v-if="visibleStudents.length === 0">
                  <td colspan="6" class="empty">No data matched your filter.</td>
                </tr>
              </tbody>
            </table>
          </div>
        </article>

        <article class="card detail-card">
          <h2>Student Detail</h2>
          <div v-if="activeStudent" class="detail-stack">
            <p class="meta">
              Active: <strong>{{ activeStudent.id }}</strong>
            </p>

            <div class="section-title">Update Basic Info</div>
            <div class="form-grid compact">
              <label>
                Name
                <input v-model="editForm.name" />
              </label>
              <label>
                Age
                <input v-model.number="editForm.age" type="number" min="1" />
              </label>
              <label>
                Gender
                <select v-model="editForm.gender">
                  <option value="male">male</option>
                  <option value="female">female</option>
                </select>
              </label>
            </div>
            <button class="btn primary" :disabled="saving" @click="handleUpdateStudent">
              Update Student
            </button>

            <div class="section-title">Update Score</div>
            <div class="form-grid compact">
              <label>
                Student
                <select v-model="scoreForm.studentId">
                  <option value="">select one</option>
                  <option v-for="item in students" :key="item.id" :value="item.id">
                    {{ item.id }} - {{ item.name }}
                  </option>
                </select>
              </label>
              <label>
                Subject
                <input v-model="scoreForm.subject" placeholder="e.g. Math" />
              </label>
              <label>
                Score
                <input v-model.number="scoreForm.score" type="number" min="0" max="100" />
              </label>
            </div>
            <button class="btn ghost" :disabled="saving" @click="handleUpdateScore">
              Update Score
            </button>

            <div class="section-title">Scores</div>
            <ul class="scores" v-if="Object.keys(activeStudent.scores || {}).length > 0">
              <li v-for="(value, subject) in activeStudent.scores" :key="subject">
                <span>{{ subject }}</span>
                <strong :class="{ fail: value < 60 }">{{ value.toFixed(1) }}</strong>
              </li>
            </ul>
            <p v-else class="empty-inline">No scores yet.</p>
          </div>

          <p v-else class="empty-inline">Pick a student from the table to manage details.</p>
        </article>
      </section>
    </main>
  </div>
</template>

<style scoped>
:global(body) {
  margin: 0;
  font-family: "Space Grotesk", "Segoe UI", "Helvetica Neue", sans-serif;
  background: #f4efe6;
}

* {
  box-sizing: border-box;
}

.page {
  --ink: #1f2a2c;
  --paper: #fffaf2;
  --line: #ddc9a7;
  --accent: #e6632e;
  --accent-2: #1f8a77;
  min-height: 100vh;
  padding: 36px 18px;
  position: relative;
  overflow: hidden;
  background:
    radial-gradient(circle at 10% 15%, #ffd9aa 0%, transparent 42%),
    radial-gradient(circle at 85% 20%, #bbf0e5 0%, transparent 36%),
    linear-gradient(135deg, #f9f2e6 0%, #f6ecdc 45%, #f8f0e3 100%);
}

.glow {
  position: absolute;
  border-radius: 999px;
  filter: blur(38px);
  opacity: 0.3;
  pointer-events: none;
}

.glow-a {
  width: 320px;
  height: 320px;
  background: #ff9c69;
  top: -80px;
  right: -100px;
}

.glow-b {
  width: 280px;
  height: 280px;
  background: #2fb9a0;
  bottom: -70px;
  left: -80px;
}

.board {
  position: relative;
  z-index: 2;
  max-width: 1180px;
  margin: 0 auto;
  display: grid;
  gap: 18px;
}

.hero {
  background: color-mix(in srgb, var(--paper) 88%, #fff 12%);
  border: 1px solid var(--line);
  border-radius: 18px;
  padding: 20px 22px;
  box-shadow: 0 8px 32px rgba(124, 86, 39, 0.12);
  animation: rise 480ms ease-out;
}

.tag {
  margin: 0;
  font-size: 12px;
  text-transform: uppercase;
  letter-spacing: 0.14em;
  color: var(--accent-2);
}

h1 {
  margin: 6px 0 8px;
  color: var(--ink);
  font-size: clamp(1.6rem, 2.4vw, 2.4rem);
}

.sub {
  margin: 0;
  color: #4f5f62;
}

.notice {
  margin: 0;
  padding: 11px 14px;
  border-radius: 10px;
  font-size: 14px;
  border: 1px solid transparent;
}

.notice.success {
  background: #ddf6ee;
  border-color: #8bd9c8;
  color: #1d6b5d;
}

.notice.error {
  background: #ffe6dc;
  border-color: #ffb79a;
  color: #8f3212;
}

.grid {
  display: grid;
  gap: 16px;
  grid-template-columns: 1fr 1fr;
}

.lower {
  grid-template-columns: 1.2fr 0.8fr;
}

.card {
  background: color-mix(in srgb, var(--paper) 90%, #fff 10%);
  border: 1px solid var(--line);
  border-radius: 16px;
  padding: 16px;
  box-shadow: 0 6px 20px rgba(95, 67, 32, 0.08);
  animation: rise 560ms ease-out;
}

.card h2 {
  margin: 0 0 12px;
  color: var(--ink);
  font-size: 1.05rem;
}

.form-grid {
  display: grid;
  gap: 10px;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  margin-bottom: 12px;
}

.form-grid.compact {
  grid-template-columns: 1fr;
}

label {
  display: grid;
  gap: 6px;
  font-size: 12px;
  text-transform: uppercase;
  letter-spacing: 0.06em;
  color: #506063;
}

input,
select,
button {
  font: inherit;
}

input,
select {
  border: 1px solid #cdb794;
  background: #fff;
  border-radius: 10px;
  padding: 9px 11px;
  color: #233437;
}

input:focus,
select:focus {
  outline: 2px solid color-mix(in srgb, var(--accent-2) 50%, white 50%);
  border-color: var(--accent-2);
}

.btn {
  border: none;
  border-radius: 10px;
  padding: 9px 14px;
  cursor: pointer;
  transition: transform 120ms ease, filter 120ms ease;
}

.btn:disabled {
  cursor: not-allowed;
  opacity: 0.65;
}

.btn:hover:not(:disabled) {
  transform: translateY(-1px);
  filter: brightness(1.02);
}

.btn.primary {
  background: linear-gradient(135deg, #ef6e39, #d7541f);
  color: #fff;
}

.btn.ghost {
  background: #e6f6f2;
  color: #1e6e5f;
  border: 1px solid #9ed8ca;
}

.btn.mini {
  padding: 6px 10px;
  font-size: 12px;
  background: #e9f2ff;
  color: #215a8f;
  border: 1px solid #bdd7f4;
}

.btn.mini.danger {
  background: #ffebe3;
  color: #8a3014;
  border-color: #f8baa2;
}

.head-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.badge {
  background: #153f39;
  color: #d4fff5;
  font-size: 12px;
  padding: 4px 10px;
  border-radius: 999px;
}

.table-wrap {
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th,
td {
  border-bottom: 1px solid #ead8b9;
  padding: 10px 8px;
  text-align: left;
  font-size: 14px;
}

th {
  color: #486164;
  font-weight: 700;
}

tr.active {
  background: #fff4e6;
}

.actions {
  display: flex;
  gap: 6px;
}

.empty {
  text-align: center;
  color: #74858a;
}

.detail-stack {
  display: grid;
  gap: 10px;
}

.meta {
  margin: 0;
  color: #415459;
}

.section-title {
  margin-top: 6px;
  font-weight: 700;
  color: #263c40;
  font-size: 13px;
  text-transform: uppercase;
  letter-spacing: 0.06em;
}

.scores {
  margin: 0;
  padding: 0;
  list-style: none;
  display: grid;
  gap: 6px;
}

.scores li {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #fef9ef;
  border: 1px solid #efdfc2;
  border-radius: 10px;
  padding: 8px 10px;
}

.fail {
  color: #ad3310;
}

.empty-inline {
  margin: 0;
  color: #708084;
}

@keyframes rise {
  from {
    opacity: 0;
    transform: translateY(8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (max-width: 960px) {
  .grid,
  .lower {
    grid-template-columns: 1fr;
  }

  .form-grid {
    grid-template-columns: 1fr;
  }

  .page {
    padding: 22px 12px;
  }
}
</style>
