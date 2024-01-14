package infra

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	errors "go-api-echo/internal/pkg/helpers/helpers_errors"
	"go-api-echo/internal/services/video/adapter"
	"go-api-echo/internal/services/video/entity"
)

type VideoProjectRepo struct {
	ctx context.Context
	db  *sqlx.DB
}

func (r VideoProjectRepo) CreateVideoProject(userID int, req *adapter.GenerateVideoReq) (*entity.VideoProject, *errors.Error) {
	res, err := r.db.ExecContext(
		r.ctx,
		`
		INSERT INTO video_project (
			user_id,
		    product_title,
		    brand_name,
			product_type,
		    market_target,
		    superiority,
		    duration
		) VALUES (?, ?, ?, ?, ?, ?, ?);
		`,
		userID,
		req.ProductTitle,
		req.BrandName,
		req.ProductType,
		req.MarketTarget,
		req.Superiority,
		req.Duration,
	)
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError("error creating video project")

		return nil, &sqlErr
	}

	projectID, err := res.LastInsertId()
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError("error getting last insert id")

		return nil, &sqlErr
	}

	createdProject, errs := r.GetVideoProject(int(projectID), userID)
	if errs != nil {
		return nil, errs
	}

	return createdProject, nil
}

func (r VideoProjectRepo) GetVideoProject(projectID, userID int) (*entity.VideoProject, *errors.Error) {
	project := entity.VideoProject{}
	err := r.db.GetContext(
		r.ctx,
		&project,
		`
		SELECT
			id,
			user_id,
			product_title,
			brand_name,
			product_type,
			market_target,
			superiority,
			duration
		FROM video_project
		WHERE id = ? AND user_id = ?
		`,
		projectID,
		userID,
	)
	if err != nil {
		fmt.Println(err)
		sqlErr := errors.FromSql(err)
		sqlErr.AddError("error getting video project")

		return nil, &sqlErr
	}

	return &project, nil
}

func (r VideoProjectRepo) GetVideoProjectList(userID int) (*[]entity.VideoProject, *errors.Error) {
	var projects []entity.VideoProject
	err := r.db.SelectContext(
		r.ctx,
		&projects,
		`
		SELECT
			id,
			user_id,
			product_title,
			brand_name,
			product_type,
			market_target,
			superiority,
			duration
		FROM video_project
		WHERE user_id = ?
		`,
		userID,
	)
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError("error getting video project list")

		return nil, &sqlErr
	}

	return &projects, nil
}

func (r VideoProjectRepo) GetSceneList(projectID int) (*[]entity.Scene, *errors.Error) {
	var scenes []entity.Scene
	err := r.db.SelectContext(
		r.ctx,
		&scenes,
		`
		SELECT
			id,
			video_project_id,
			sequence,
			title,
			narration,
			illustration,
			illustration_url,
			voice_url
		FROM scene
		WHERE video_project_id = ?
		`,
		projectID,
	)

	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError("error getting scene list")

		return nil, &sqlErr
	}

	return &scenes, nil
}

func (r VideoProjectRepo) GetVideoProjectDetail(projectID int, userID int) (*adapter.VideoProjectDetail, *errors.Error) {
	project, err := r.GetVideoProject(projectID, userID)
	if err != nil {
		return nil, err
	}

	scenes, err := r.GetSceneList(projectID)
	if err != nil {
		return nil, err
	}

	return &adapter.VideoProjectDetail{
		VideoProject: *project,
		Scenes:       *scenes,
	}, nil
}

func (r VideoProjectRepo) CreateScene(projectID int, scene *entity.Scene) (int, *errors.Error) {
	res, err := r.db.ExecContext(
		r.ctx,
		`
		INSERT INTO scene (
			video_project_id,
			sequence,
			title,
			narration,
			illustration,
			illustration_url,
			voice_url
		) VALUES (?, ?, ?, ?, ?, ?, ?);
		`,
		projectID,
		scene.Sequence,
		scene.Title,
		scene.Narration,
		scene.Illustration,
		scene.IllustrationUrl,
		scene.VoiceUrl,
	)
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError("error creating scene")

		return 0, &sqlErr
	}

	sceneID, err := res.LastInsertId()
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError("error getting last insert id")

		return 0, &sqlErr
	}

	return int(sceneID), nil
}

func (r VideoProjectRepo) UpdateScene(sceneID int, scene *entity.Scene) (int, *errors.Error) {
	res, err := r.db.ExecContext(
		r.ctx,
		`
		UPDATE scene SET
			illustration_url = COALESCE(?, illustration_url),
			voice_url = COALESCE(?, voice_url)
		WHERE id = ?
		`,
		&scene.IllustrationUrl,
		&scene.VoiceUrl,
		sceneID,
	)
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError("error updating scene")

		return 0, &sqlErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError("error getting rows affected")

		return 0, &sqlErr
	}

	return int(rowsAffected), nil
}
