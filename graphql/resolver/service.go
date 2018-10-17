package resolver

import (
	"context"
	"fmt"

	"github.com/hsukvn/go-gin-graphql-template/graphql/model"
	"github.com/hsukvn/go-gin-graphql-template/graphql/scalar"
	"github.com/hsukvn/go-gin-graphql-template/lib/unit"
)

type serviceArgs struct {
	Name string
}

func (r *Resolver) Service(ctx context.Context, args serviceArgs) (*serviceResolver, error) {
	s, err := unit.NewService(args.Name)
	if err != nil {
		return nil, fmt.Errorf("service: Fail to new service (%v), err: (%v)", args.Name, err)
	}

	service, err := model.NewService(s)
	if err != nil {
		return nil, fmt.Errorf("service: Fail to new service model (%v), err: (%v)", args.Name, err)
	}

	return &serviceResolver{service: service}, nil
}

func (r *Resolver) StartService(ctx context.Context, args serviceArgs) (*serviceResolver, error) {
	s, err := unit.NewService(args.Name)
	if err != nil {
		return nil, fmt.Errorf("service: Fail to new service (%v), err: (%v)", args.Name, err)
	}

	err = s.Start()
	if err != nil {
		return nil, fmt.Errorf("service: Fail to start service (%v), err: (%v)", args.Name, err)
	}

	service, err := model.NewService(s)
	if err != nil {
		return nil, fmt.Errorf("service: Fail to new service model (%v), err: (%v)", args.Name, err)
	}

	return &serviceResolver{service: service}, nil
}

func (r *Resolver) StopService(ctx context.Context, args serviceArgs) (*serviceResolver, error) {
	s, err := unit.NewService(args.Name)
	if err != nil {
		return nil, fmt.Errorf("service: Fail to new service (%v), err: (%v)", args.Name, err)
	}

	err = s.Stop()
	if err != nil {
		return nil, fmt.Errorf("service: Fail to stop service (%v), err: (%v)", args.Name, err)
	}

	service, err := model.NewService(s)
	if err != nil {
		return nil, fmt.Errorf("service: Fail to new service model (%v), err: (%v)", args.Name, err)
	}

	return &serviceResolver{service: service}, nil
}

func (r *Resolver) EnableService(ctx context.Context, args serviceArgs) (*serviceResolver, error) {
	s, err := unit.NewService(args.Name)
	if err != nil {
		return nil, fmt.Errorf("service: Fail to new service (%v), err: (%v)", args.Name, err)
	}

	err = s.Enable()
	if err != nil {
		return nil, fmt.Errorf("service: Fail to enable service (%v), err: (%v)", args.Name, err)
	}

	service, err := model.NewService(s)
	if err != nil {
		return nil, fmt.Errorf("service: Fail to new service model (%v), err: (%v)", args.Name, err)
	}

	return &serviceResolver{service: service}, nil
}

func (r *Resolver) DisableService(ctx context.Context, args serviceArgs) (*serviceResolver, error) {
	s, err := unit.NewService(args.Name)
	if err != nil {
		return nil, fmt.Errorf("service: Fail to new service (%v), err: (%v)", args.Name, err)
	}

	err = s.Disable()
	if err != nil {
		return nil, fmt.Errorf("service: Fail to disable service (%v), err: (%v)", args.Name, err)
	}

	service, err := model.NewService(s)
	if err != nil {
		return nil, fmt.Errorf("service: Fail to new service model (%v), err: (%v)", args.Name, err)
	}

	return &serviceResolver{service: service}, nil
}

type serviceResolver struct {
	service *model.Service
}

func (r *serviceResolver) Name() *string {
	return &r.service.Name
}

func (r *serviceResolver) MainPID() *scalar.Uint64 {
	pid := scalar.Uint64(r.service.MainPID)
	return &pid
}

func (r *serviceResolver) ActiveState() *int32 {
	return &r.service.ActiveState
}

func (r *serviceResolver) UnitFileState() *int32 {
	return &r.service.UnitFileState
}
