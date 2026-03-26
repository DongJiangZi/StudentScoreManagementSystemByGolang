import request from "./request";
import type {
  Student,
  CreateStudentDTO,
  UpdateStudentDTO,
  UpdateScoreDTO,
  ApiResponse,
} from "@/types/student";

export function getStudentList() {
  return request.get<any, ApiResponse<Student[]>>("/students");
}

export function getStudentDetail(id: string) {
  return request.get<any, ApiResponse<Student>>(`/students/${id}`);
}

export function createStudent(data: CreateStudentDTO) {
  return request.post<any, ApiResponse<Student>>("/students", data);
}

export function updateStudent(id: string, data: UpdateStudentDTO) {
  return request.put<any, ApiResponse<null>>(`/students/${id}`, data);
}

export function deleteStudent(id: string) {
  return request.delete<any, ApiResponse<null>>(`/students/${id}`);
}

export function updateStudentScore(
  id: string,
  subject: string,
  data: UpdateScoreDTO,
) {
  return request.put<any, ApiResponse<null>>(
    `/students/${id}/scores/${subject}`,
    data,
  );
}
