# EduCore Student Management System

一个用于巩固 **Go 基础语法** 与 **前后端交互** 的练习项目。

项目目标是通过一个完整但轻量的学生管理系统，把 Go 的核心能力（分层设计、错误处理、日志、文件存储、HTTP 接口）和前端调用流程串起来，为后续学习 Go 框架（如 Gin、Fiber、Echo）打基础。

## 学习目标

- 巩固 Go 基础语法与工程组织方式
- 熟悉分层架构（domain / repository / service / interfaces）
- 掌握错误处理与日志记录的基本实践
- 理解前后端接口联调、数据结构对齐、CORS 等交互要点
- 为后续迁移到 Go Web 框架建立可复用的业务层基础

## 技术栈

- 后端：Go 1.26
- 前端：Vue 3 + TypeScript + Vite + Axios
- 存储：内存仓储 + JSON 文件仓储

## 项目结构

```text
EduCoreStudentManagementSystem/
├─ backend/
│  ├─ cmd/
│  │  ├─ cli/      # 命令行入口
│  │  └─ http/     # HTTP 服务入口
│  ├─ internal/
│  │  ├─ domain/   # 领域模型
│  │  ├─ repository/ # 仓储接口与实现（memory/json）
│  │  ├─ app/service/ # 业务服务层
│  │  ├─ interfaces/  # CLI / HTTP 适配层
│  │  └─ infrastructure/ # 配置、工厂、日志
│  ├─ data/        # JSON 数据文件
│  └─ logs/        # 日志文件
└─ frontend/
   └─ src/
      ├─ api/      # 请求封装与接口方法
      ├─ views/    # 页面
      ├─ router/   # 路由
      └─ types/    # 类型定义
```

## 已实现功能

- 学生管理基础 CRUD
- 查询学生详情
- 更新学生基本信息
- 更新学生单科成绩
- 学生列表展示与删除
- 前端筛选与排行展示（关键字、性别、失败科目、TopN）
- 后端统一错误码响应结构
- 文件日志与异步日志写入

## 快速启动

### 1. 启动后端（HTTP）

```bash
cd backend
go run ./cmd/http
```

默认监听：`http://localhost:8080`

### 2. 启动前端

```bash
cd frontend
npm install
npm run dev
```

默认地址：`http://localhost:5173`

## 常用命令

### 后端

```bash
cd backend
go test ./...
go build ./...
```

### 前端

```bash
cd frontend
npm run build
```

## 当前学习重点与下一步计划

- 继续完善接口测试与异常场景覆盖
- 补充鉴权、参数校验与中间件思路
- 将现有 service/repository 迁移到 Go 框架（优先 Gin）
- 引入数据库（MySQL/PostgreSQL）替代 JSON 持久化
- 逐步加入更完整的工程化能力（配置分环境、日志分级、部署脚本）

## 说明

这是一个学习驱动的项目，重点在于“理解过程”和“可持续迭代”，而不是一次性做成复杂系统。

如果你正在学习 Go，这个项目可以作为从基础语法过渡到框架开发的中间练习样板。
