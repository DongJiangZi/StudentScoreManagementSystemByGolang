export interface Student {
  id: string;
  name: string;
  age: number;
  gender: string;
  scores: Record<string, number>;
  createdAt?: string;
  updatedAt?: string;
}

export interface CreateStudentDTO {
  id: string;
  name: string;
  age: number;
  gender: string;
}

export interface UpdateStudentDTO {
  name: string;
  age: number;
  gender: string;
}

export interface UpdateScoreDTO {
  score: number;
}

export interface ApiResponse<T> {
  code: number;
  message: string;
  data: T;
}
