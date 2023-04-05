// Code generated by goctl. DO NOT EDIT.
// Source: project.proto

package server

import (
	"context"

	"mathless-backend/app/project/cmd/rpc/internal/logic"
	"mathless-backend/app/project/cmd/rpc/internal/svc"
	"mathless-backend/app/project/cmd/rpc/project"
)

type ProjectServiceServer struct {
	svcCtx *svc.ServiceContext
	project.UnimplementedProjectServiceServer
}

func NewProjectServiceServer(svcCtx *svc.ServiceContext) *ProjectServiceServer {
	return &ProjectServiceServer{
		svcCtx: svcCtx,
	}
}

// 获取用户的工程
func (s *ProjectServiceServer) GetProjects(ctx context.Context, in *project.GetProjectsRequest) (*project.GetProjectsResponse, error) {
	l := logic.NewGetProjectsLogic(ctx, s.svcCtx)
	return l.GetProjects(in)
}

// 创建工程
func (s *ProjectServiceServer) CreateProject(ctx context.Context, in *project.CreateProjectRequest) (*project.CreateProjectResponse, error) {
	l := logic.NewCreateProjectLogic(ctx, s.svcCtx)
	return l.CreateProject(in)
}

// 删除工程
func (s *ProjectServiceServer) DeleteProject(ctx context.Context, in *project.DeleteProjectRequest) (*project.DeleteProjectResponse, error) {
	l := logic.NewDeleteProjectLogic(ctx, s.svcCtx)
	return l.DeleteProject(in)
}

// 更新工程
func (s *ProjectServiceServer) UpdateProject(ctx context.Context, in *project.UpdateProjectRequest) (*project.UpdateProjectResponse, error) {
	l := logic.NewUpdateProjectLogic(ctx, s.svcCtx)
	return l.UpdateProject(in)
}
