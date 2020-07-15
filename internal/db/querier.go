// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateLabelColor(ctx context.Context, arg CreateLabelColorParams) (LabelColor, error)
	CreateOrganization(ctx context.Context, arg CreateOrganizationParams) (Organization, error)
	CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error)
	CreateProjectLabel(ctx context.Context, arg CreateProjectLabelParams) (ProjectLabel, error)
	CreateProjectMember(ctx context.Context, arg CreateProjectMemberParams) (ProjectMember, error)
	CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) (RefreshToken, error)
	CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error)
	CreateTaskAssigned(ctx context.Context, arg CreateTaskAssignedParams) (TaskAssigned, error)
	CreateTaskChecklist(ctx context.Context, arg CreateTaskChecklistParams) (TaskChecklist, error)
	CreateTaskChecklistItem(ctx context.Context, arg CreateTaskChecklistItemParams) (TaskChecklistItem, error)
	CreateTaskGroup(ctx context.Context, arg CreateTaskGroupParams) (TaskGroup, error)
	CreateTaskLabelForTask(ctx context.Context, arg CreateTaskLabelForTaskParams) (TaskLabel, error)
	CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error)
	CreateTeamMember(ctx context.Context, arg CreateTeamMemberParams) (TeamMember, error)
	CreateUserAccount(ctx context.Context, arg CreateUserAccountParams) (UserAccount, error)
	DeleteExpiredTokens(ctx context.Context) error
	DeleteProjectByID(ctx context.Context, projectID uuid.UUID) error
	DeleteProjectLabelByID(ctx context.Context, projectLabelID uuid.UUID) error
	DeleteProjectMember(ctx context.Context, arg DeleteProjectMemberParams) error
	DeleteRefreshTokenByID(ctx context.Context, tokenID uuid.UUID) error
	DeleteRefreshTokenByUserID(ctx context.Context, userID uuid.UUID) error
	DeleteTaskAssignedByID(ctx context.Context, arg DeleteTaskAssignedByIDParams) (TaskAssigned, error)
	DeleteTaskByID(ctx context.Context, taskID uuid.UUID) error
	DeleteTaskChecklistByID(ctx context.Context, taskChecklistID uuid.UUID) error
	DeleteTaskChecklistItem(ctx context.Context, taskChecklistItemID uuid.UUID) error
	DeleteTaskGroupByID(ctx context.Context, taskGroupID uuid.UUID) (int64, error)
	DeleteTaskLabelByID(ctx context.Context, taskLabelID uuid.UUID) error
	DeleteTaskLabelForTaskByProjectLabelID(ctx context.Context, arg DeleteTaskLabelForTaskByProjectLabelIDParams) error
	DeleteTasksByTaskGroupID(ctx context.Context, taskGroupID uuid.UUID) (int64, error)
	DeleteTeamByID(ctx context.Context, teamID uuid.UUID) error
	DeleteTeamMember(ctx context.Context, arg DeleteTeamMemberParams) error
	DeleteUserAccountByID(ctx context.Context, userID uuid.UUID) error
	GetAllOrganizations(ctx context.Context) ([]Organization, error)
	GetAllProjects(ctx context.Context) ([]Project, error)
	GetAllProjectsForTeam(ctx context.Context, teamID uuid.UUID) ([]Project, error)
	GetAllTaskGroups(ctx context.Context) ([]TaskGroup, error)
	GetAllTasks(ctx context.Context) ([]Task, error)
	GetAllTeams(ctx context.Context) ([]Team, error)
	GetAllUserAccounts(ctx context.Context) ([]UserAccount, error)
	GetAssignedMembersForTask(ctx context.Context, taskID uuid.UUID) ([]TaskAssigned, error)
	GetLabelColorByID(ctx context.Context, labelColorID uuid.UUID) (LabelColor, error)
	GetLabelColors(ctx context.Context) ([]LabelColor, error)
	GetOwnedTeamProjectsForUserID(ctx context.Context, arg GetOwnedTeamProjectsForUserIDParams) ([]uuid.UUID, error)
	GetProjectByID(ctx context.Context, projectID uuid.UUID) (Project, error)
	GetProjectIDForTask(ctx context.Context, taskID uuid.UUID) (uuid.UUID, error)
	GetProjectLabelByID(ctx context.Context, projectLabelID uuid.UUID) (ProjectLabel, error)
	GetProjectLabelsForProject(ctx context.Context, projectID uuid.UUID) ([]ProjectLabel, error)
	GetProjectMembersForProjectID(ctx context.Context, projectID uuid.UUID) ([]ProjectMember, error)
	GetRefreshTokenByID(ctx context.Context, tokenID uuid.UUID) (RefreshToken, error)
	GetRoleForProjectMemberByUserID(ctx context.Context, arg GetRoleForProjectMemberByUserIDParams) (Role, error)
	GetRoleForTeamMember(ctx context.Context, arg GetRoleForTeamMemberParams) (Role, error)
	GetRoleForUserID(ctx context.Context, userID uuid.UUID) (GetRoleForUserIDRow, error)
	GetTaskByID(ctx context.Context, taskID uuid.UUID) (Task, error)
	GetTaskChecklistByID(ctx context.Context, taskChecklistID uuid.UUID) (TaskChecklist, error)
	GetTaskChecklistItemByID(ctx context.Context, taskChecklistItemID uuid.UUID) (TaskChecklistItem, error)
	GetTaskChecklistItemsForTaskChecklist(ctx context.Context, taskChecklistID uuid.UUID) ([]TaskChecklistItem, error)
	GetTaskChecklistsForTask(ctx context.Context, taskID uuid.UUID) ([]TaskChecklist, error)
	GetTaskGroupByID(ctx context.Context, taskGroupID uuid.UUID) (TaskGroup, error)
	GetTaskGroupsForProject(ctx context.Context, projectID uuid.UUID) ([]TaskGroup, error)
	GetTaskLabelByID(ctx context.Context, taskLabelID uuid.UUID) (TaskLabel, error)
	GetTaskLabelForTaskByProjectLabelID(ctx context.Context, arg GetTaskLabelForTaskByProjectLabelIDParams) (TaskLabel, error)
	GetTaskLabelsForTaskID(ctx context.Context, taskID uuid.UUID) ([]TaskLabel, error)
	GetTasksForTaskGroupID(ctx context.Context, taskGroupID uuid.UUID) ([]Task, error)
	GetTeamByID(ctx context.Context, teamID uuid.UUID) (Team, error)
	GetTeamMemberByID(ctx context.Context, arg GetTeamMemberByIDParams) (TeamMember, error)
	GetTeamMembersForTeamID(ctx context.Context, teamID uuid.UUID) ([]TeamMember, error)
	GetTeamsForOrganization(ctx context.Context, organizationID uuid.UUID) ([]Team, error)
	GetUserAccountByID(ctx context.Context, userID uuid.UUID) (UserAccount, error)
	GetUserAccountByUsername(ctx context.Context, username string) (UserAccount, error)
	SetProjectOwner(ctx context.Context, arg SetProjectOwnerParams) (Project, error)
	SetTaskChecklistItemComplete(ctx context.Context, arg SetTaskChecklistItemCompleteParams) (TaskChecklistItem, error)
	SetTaskComplete(ctx context.Context, arg SetTaskCompleteParams) (Task, error)
	SetTaskGroupName(ctx context.Context, arg SetTaskGroupNameParams) (TaskGroup, error)
	SetTeamOwner(ctx context.Context, arg SetTeamOwnerParams) (Team, error)
	SetUserPassword(ctx context.Context, arg SetUserPasswordParams) (UserAccount, error)
	UpdateProjectLabel(ctx context.Context, arg UpdateProjectLabelParams) (ProjectLabel, error)
	UpdateProjectLabelColor(ctx context.Context, arg UpdateProjectLabelColorParams) (ProjectLabel, error)
	UpdateProjectLabelName(ctx context.Context, arg UpdateProjectLabelNameParams) (ProjectLabel, error)
	UpdateProjectMemberRole(ctx context.Context, arg UpdateProjectMemberRoleParams) (ProjectMember, error)
	UpdateProjectNameByID(ctx context.Context, arg UpdateProjectNameByIDParams) (Project, error)
	UpdateTaskChecklistItemLocation(ctx context.Context, arg UpdateTaskChecklistItemLocationParams) (TaskChecklistItem, error)
	UpdateTaskChecklistItemName(ctx context.Context, arg UpdateTaskChecklistItemNameParams) (TaskChecklistItem, error)
	UpdateTaskChecklistName(ctx context.Context, arg UpdateTaskChecklistNameParams) (TaskChecklist, error)
	UpdateTaskChecklistPosition(ctx context.Context, arg UpdateTaskChecklistPositionParams) (TaskChecklist, error)
	UpdateTaskDescription(ctx context.Context, arg UpdateTaskDescriptionParams) (Task, error)
	UpdateTaskDueDate(ctx context.Context, arg UpdateTaskDueDateParams) (Task, error)
	UpdateTaskGroupLocation(ctx context.Context, arg UpdateTaskGroupLocationParams) (TaskGroup, error)
	UpdateTaskLocation(ctx context.Context, arg UpdateTaskLocationParams) (Task, error)
	UpdateTaskName(ctx context.Context, arg UpdateTaskNameParams) (Task, error)
	UpdateTeamMemberRole(ctx context.Context, arg UpdateTeamMemberRoleParams) (TeamMember, error)
	UpdateUserAccountProfileAvatarURL(ctx context.Context, arg UpdateUserAccountProfileAvatarURLParams) (UserAccount, error)
	UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) (UserAccount, error)
}

var _ Querier = (*Queries)(nil)
