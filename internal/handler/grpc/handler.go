package grpc

import (
	"context"
	"strings"

	"github.com/egorgasay/gost"
	api "github.com/egorgasay/itisadb-shared-proto/go"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"itisadb/config"
	"itisadb/internal/constants"
	"itisadb/internal/domains"
	"itisadb/internal/handler/converterr"
	"itisadb/internal/models"
)

type Handler struct {
	api.UnimplementedItisaDBServer
	core     domains.Balancer
	logger   *zap.Logger
	session  domains.Session
	security config.SecurityConfig
	converterr converterr.ConvertErr
}

func New(
	logic domains.Balancer,
	l *zap.Logger,
	session domains.Session,
	conf config.SecurityConfig,
	converterr converterr.ConvertErr,
) *Handler {
	return &Handler{core: logic, logger: l, session: session, security: conf, converterr: converterr}
}

func (h *Handler) claimsFromContext(ctx context.Context) (opt gost.Option[models.UserClaims]) {
	value := ctx.Value(constants.UserKey)
	if value == nil {
		return opt.None()
	}

	claims, ok := value.(models.UserClaims)
	if !ok {
		h.logger.Warn("failed to cast userID", zap.Any("value", value))
		return opt.None()
	}

	return opt.Some(claims)
}

func (h *Handler) Set(ctx context.Context, r *api.SetRequest) (*api.SetResponse, error) {
	claims := h.claimsFromContext(ctx)

	opts := models.SetOptions{}
	if err := copier.Copy(&opts, gost.SafeDeref(r.Options)); err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	setTo, err := h.core.Set(ctx, claims, r.Key, r.Value, opts)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.SetResponse{
		SavedTo: setTo,
	}, nil
}

func (h *Handler) SetToObject(ctx context.Context, r *api.SetToObjectRequest) (*api.SetToObjectResponse, error) {
	claims := h.claimsFromContext(ctx)

	opts := models.SetToObjectOptions{}
	err := copier.Copy(&opts, gost.SafeDeref(r.Options))
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	setTo, err := h.core.SetToObject(ctx, claims, r.Object, r.Key, r.Value, opts)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.SetToObjectResponse{
		SavedTo: setTo,
	}, nil
}

func (h *Handler) Get(ctx context.Context, r *api.GetRequest) (*api.GetResponse, error) {
	claims := h.claimsFromContext(ctx)

	opts := models.GetOptions{}
	err := copier.Copy(&opts, gost.SafeDeref(r.Options))
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	res, err := h.core.Get(ctx, claims, r.Key, opts)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.GetResponse{
		Value:    res.Value,
		ReadOnly: res.ReadOnly,
		Level:    api.Level(res.Level),
	}, nil
}

func (h *Handler) GetFromObject(ctx context.Context, r *api.GetFromObjectRequest) (*api.GetFromObjectResponse, error) {
	claims := h.claimsFromContext(ctx)

	opts := models.GetFromObjectOptions{}
	err := copier.Copy(&opts, gost.SafeDeref(r.Options))
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	value, err := h.core.GetFromObject(ctx, claims, r.Object, r.Key, opts)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.GetFromObjectResponse{
		Value: value,
	}, nil
}

func (h *Handler) Delete(ctx context.Context, r *api.DeleteRequest) (*api.DeleteResponse, error) {
	claims := h.claimsFromContext(ctx)

	opts := models.DeleteOptions{}
	err := copier.Copy(&opts, gost.SafeDeref(r.Options))
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	err = h.core.Delete(ctx, claims, r.Key, opts)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.DeleteResponse{}, nil
}

func (h *Handler) AttachToObject(ctx context.Context, r *api.AttachToObjectRequest) (*api.AttachToObjectResponse, error) {
	claims := h.claimsFromContext(ctx)

	opts := models.AttachToObjectOptions{}
	err := copier.Copy(&opts, gost.SafeDeref(r.Options))
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	err = h.core.AttachToObject(ctx, claims, r.Dst, r.Src, opts)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.AttachToObjectResponse{}, nil
}

func (h *Handler) DeleteObject(ctx context.Context, r *api.DeleteObjectRequest) (*api.DeleteObjectResponse, error) {
	claims := h.claimsFromContext(ctx)

	opts := models.DeleteObjectOptions{}
	err := copier.Copy(&opts, gost.SafeDeref(r.Options))
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	err = h.core.DeleteObject(ctx, claims, r.Object, opts)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.DeleteObjectResponse{}, nil
}

func (h *Handler) Connect(ctx context.Context, request *api.ConnectRequest) (*api.ConnectResponse, error) {
	serverNum, err := h.core.Connect(ctx, request.GetAddress())
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.ConnectResponse{
		Status: "connected successfully",
		Server: serverNum,
	}, nil
}

func (h *Handler) Object(ctx context.Context, r *api.ObjectRequest) (*api.ObjectResponse, error) {
	claims := h.claimsFromContext(ctx)

	opts := models.ObjectOptions{}
	err := copier.Copy(&opts, gost.SafeDeref(r.Options))
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	serv, err := h.core.Object(ctx, claims, r.Name, opts)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.ObjectResponse{
		Server: serv,
	}, nil
}

func (h *Handler) ObjectToJSON(ctx context.Context, r *api.ObjectToJSONRequest) (*api.ObjectToJSONResponse, error) {
	claims := h.claimsFromContext(ctx)

	opts := models.ObjectToJSONOptions{}
	err := copier.Copy(&opts, gost.SafeDeref(r.Options))
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	m, err := h.core.ObjectToJSON(ctx, claims, r.Name, opts)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.ObjectToJSONResponse{
		Object: m,
	}, nil
}

func (h *Handler) IsObject(ctx context.Context, r *api.IsObjectRequest) (*api.IsObjectResponse, error) {
	claims := h.claimsFromContext(ctx)

	opts := models.IsObjectOptions{}
	err := copier.Copy(&opts, gost.SafeDeref(r.Options))
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	ok, err := h.core.IsObject(ctx, claims, r.Name, opts)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.IsObjectResponse{
		Ok: ok,
	}, nil
}

func (h *Handler) DeleteAttr(ctx context.Context, r *api.DeleteAttrRequest) (*api.DeleteAttrResponse, error) {
	claims := h.claimsFromContext(ctx)

	opts := models.DeleteAttrOptions{}
	err := copier.Copy(&opts, gost.SafeDeref(r.Options))
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	err = h.core.DeleteAttr(ctx, claims, r.Key, r.Object, opts)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.DeleteAttrResponse{}, nil
}

func (h *Handler) Size(ctx context.Context, r *api.ObjectSizeRequest) (*api.ObjectSizeResponse, error) {
	claims := h.claimsFromContext(ctx)

	opts := models.SizeOptions{}
	err := copier.Copy(&opts, gost.SafeDeref(r.Options))
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	size, err := h.core.Size(ctx, claims, r.Name, opts)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.ObjectSizeResponse{
		Size: size,
	}, nil
}

func (h *Handler) Disconnect(ctx context.Context, r *api.DisconnectRequest) (*api.DisconnectResponse, error) {
	err := h.core.Disconnect(ctx, r.Server)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.DisconnectResponse{}, nil
}

func (h *Handler) Servers(ctx context.Context, r *api.ServersRequest) (*api.ServersResponse, error) {
	servers := h.core.Servers()
	s := strings.Join(servers, "\n")
	return &api.ServersResponse{
		ServersInfo: s,
	}, nil
}

func (h *Handler) Authenticate(ctx context.Context, request *api.AuthRequest) (*api.AuthResponse, error) {
	token, err := h.core.Authenticate(ctx, request.GetLogin(), request.GetPassword())
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.AuthResponse{Token: token}, nil
}

func (h *Handler) NewUser(ctx context.Context, r *api.NewUserRequest) (*api.NewUserResponse, error) {
	claims := h.claimsFromContext(ctx)

	user := models.User{}
	err := copier.Copy(&user, gost.SafeDeref(r.User))
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	err = h.core.NewUser(ctx, claims, user)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.NewUserResponse{}, nil
}
func (h *Handler) DeleteUser(ctx context.Context, r *api.DeleteUserRequest) (*api.DeleteUserResponse, error) {
	claims := h.claimsFromContext(ctx)

	err := h.core.DeleteUser(ctx, claims, r.Login)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.DeleteUserResponse{}, nil
}
func (h *Handler) ChangePassword(ctx context.Context, r *api.ChangePasswordRequest) (*api.ChangePasswordResponse, error) {
	claims := h.claimsFromContext(ctx)

	err := h.core.ChangePassword(ctx, claims, r.Login, r.NewPassword)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.ChangePasswordResponse{}, nil
}
func (h *Handler) ChangeLevel(ctx context.Context, r *api.ChangeLevelRequest) (*api.ChangeLevelResponse, error) {
	claims := h.claimsFromContext(ctx)

	err := h.core.ChangeLevel(ctx, claims, r.Login, models.Level(r.Level))
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.ChangeLevelResponse{}, nil
}

func (h *Handler) GetRam(ctx context.Context, r *api.GetRamRequest) (*api.GetRamResponse, error) {
	res := h.core.CalculateRAM(ctx)
	if res.IsErr() {
		return nil, h.converterr.ToGRPC(res.Error())
	}

	ram := res.Unwrap()

	return &api.GetRamResponse{Ram: &api.Ram{
		Total:     ram.Total,
		Available: ram.Available,
	}}, nil
}

func apiUsersToModelUsers(usersApi []*api.User) (users []models.User) {
	for _, userApi := range usersApi {
		user := models.User{
			Login:    userApi.Login,
			Password: userApi.Password,
			Level:    models.Level(userApi.Level),
		}
		users = append(users, user)
	}

	return
}

func (h *Handler) Sync(ctx context.Context, req *api.SyncData) (*api.SyncData, error) {
	res := h.core.Sync(ctx, req.SyncID, apiUsersToModelUsers(req.Users))
	if res.IsErr() {
		return nil, h.converterr.ToGRPC(res.Error())
	}

	return req, nil
}

func (h *Handler) GetLastUserChangeID(ctx context.Context, _ *api.GetLastUserChangeIDRequest) (*api.GetLastUserChangeIDResponse, error) {
	res := h.core.GetLastUserChangeID(ctx)
	if res.IsErr() {
		return nil, h.converterr.ToGRPC(res.Error())
	}

	return &api.GetLastUserChangeIDResponse{
		LastChangeID: res.Unwrap(),
	}, nil
}

func (h *Handler) AddServer(ctx context.Context, r *api.AddServerRequest) (*api.AddServerResponse, error) {
	num, err := h.core.Connect(ctx, r.Address)
	if err != nil {
		return nil, h.converterr.ToGRPC(err)
	}

	return &api.AddServerResponse{
		ServerID: uint64(num),
	}, nil
}

