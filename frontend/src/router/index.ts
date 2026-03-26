import { createRouter, createWebHistory } from "vue-router";
import StudentView from "@/views/StudentView.vue";

const routes = [
  {
    path: "/",
    redirect: "/students",
  },
  {
    path: "/students",
    name: "Student",
    component: StudentView,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
