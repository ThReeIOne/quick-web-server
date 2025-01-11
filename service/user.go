package service

import (
	"quick_web_golang/model"
)

type UserService struct{}

//func GetSessionUid(ctx context.Context) (int, error) {
//	uid := provider.SessionManager.Manager.GetInt(ctx, lib.Uid)
//	if uid == 0 {
//		return uid, Unauthenticated
//	}
//	return uid, nil
//}

//func (s *UserService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
//	user, err := model.Repos.UserRepo.GetByUsername(in.GetUsername())
//	if err != nil {
//		_ = log.Error(err)
//		return nil, InternalError
//	}
//
//	provider.SessionManager.Manager.Put(ctx, lib.Uid, user.Id)
//	if err = provider.SessionManager.Manager.RenewToken(ctx); err != nil {
//		_ = log.Error(err)
//		return nil, InternalError
//	}
//	if _, _, err = provider.SessionManager.Manager.Commit(ctx); err != nil {
//		_ = log.Error(err)
//		return nil, InternalError
//	}
//
//	return nil, nil
//}

//func (s *UserService) Logout(ctx context.Context, _ *pb.LogoutRequest) (*pb.LogoutResponse, error) {
//	if err := provider.SessionManager.Manager.Destroy(ctx); err != nil {
//		return nil, InternalError
//	}
//	return &pb.LogoutResponse{}, nil
//}
//
//func (s *UserService) Get(ctx context.Context, _ *pb.GetRequest) (*pb.GetResponse, error) {
//	uid, _ := GetSessionUid(ctx)
//	user, err := model.Repos.UserRepo.Get(uid)
//	if err != nil {
//		_ = log.Error(err)
//		return nil, InternalError
//	}
//	return &pb.GetResponse{
//		User: &pb.User{
//			Id:       user.Id,
//			Username: user.Username,
//		},
//	}, nil
//}

func (*UserService) GetByPhone(phone string) (*model.User, error) {
	return model.Repos.UserRepo.GetByPhone(phone)
}

func (s *UserService) ExistPhone(companyId int, phone string) (bool, error) {
	return model.Repos.UserRepo.ExistPhone(companyId, phone)
}

func (s *UserService) CreateUser(user *model.User) error {

	if _user, err := model.Repos.UserRepo.GetByPhone(user.Phone); err != nil {
		return err
	} else if _user != nil {
		user.Id = _user.Id
	}

	trx, err := model.Repos.UserRepo.GetConn().Beginx()
	if err != nil {
		return err
	}

	if user.Id == 0 {
		if err = model.Repos.UserRepo.CreateUser(trx, user); err != nil {
			_ = trx.Rollback()
			return err
		}
	}

	if err = trx.Commit(); err != nil {
		_ = trx.Rollback()
		return err
	}
	return nil
}

func (s *UserService) UpdateLastLoginAt(id int) error {
	return model.Repos.UserRepo.UpdateLastLoginAt(id)
}
