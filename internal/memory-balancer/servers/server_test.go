package servers

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"itisadb/pkg/api/storage"
	storagemock "itisadb/pkg/api/storage/gomocks"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
)

func TestServer_AttachToIndex(t *testing.T) {
	type args struct {
		ctx context.Context
		dst string
		src string
	}
	tests := []struct {
		name         string
		mockBehavior func(r *storagemock.MockStorageClient)
		args         args
		wantCode     error
	}{
		{
			name: "Success",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().AttachToIndex(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			args: args{
				ctx: context.Background(),
				dst: "test",
				src: "inner",
			},
		},
		{
			name: "badConnection",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().AttachToIndex(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.Unavailable, "bad connection"))
			},
			args: args{
				ctx: context.Background(),
				dst: "test2",
				src: "inner2",
			},
			wantCode: status.Error(codes.Unavailable, "bad connection"),
		},
		{
			name: "indexNotFound",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().AttachToIndex(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.NotFound, "index not found"))
			},
			args: args{
				ctx: context.Background(),
				dst: "test3",
				src: "inner3",
			},
			wantCode: status.Error(codes.NotFound, "index not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			cl := storagemock.NewMockStorageClient(c)
			tt.mockBehavior(cl)

			s := &Server{
				tries:   atomic.Uint32{},
				storage: cl,
				ram: RAM{
					available: 100,
					total:     100,
				},
				number: 1,
				mu:     &sync.RWMutex{},
			}
			if err := s.AttachToIndex(tt.args.ctx, tt.args.dst, tt.args.src); (err != nil) && (!errors.Is(err, tt.wantCode)) {
				t.Errorf("AttachToIndex() error = %v, wantCode %v", err, tt.wantCode)
			}
		})
	}
}

func TestServer_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		Key string
	}
	tests := []struct {
		name         string
		mockBehavior func(cl *storagemock.MockStorageClient)
		args         args
		wantErr      error
	}{
		{
			name: "Success",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			args: args{
				ctx: context.Background(),
				Key: "test",
			},
		},
		{
			name: "badConnection",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.Unavailable, "bad connection"))
			},
			args: args{
				ctx: context.Background(),
				Key: "test2",
			},
			wantErr: status.Error(codes.Unavailable, "bad connection"),
		},
		{
			name: "keyNotFound",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.NotFound, "key was not found"))
			},
			args: args{
				ctx: context.Background(),
				Key: "test2",
			},
			wantErr: status.Error(codes.NotFound, "key was not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			cl := storagemock.NewMockStorageClient(c)
			tt.mockBehavior(cl)

			s := &Server{
				tries:   atomic.Uint32{},
				storage: cl,
				ram: RAM{
					available: 100,
					total:     100,
				},
				number: 1,
				mu:     &sync.RWMutex{},
			}
			if err := s.Delete(tt.args.ctx, tt.args.Key); (err != nil) && (!errors.Is(err, tt.wantErr)) {
				t.Errorf("Delete() error = %v, wantCode %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_DeleteAttr(t *testing.T) {
	type args struct {
		ctx   context.Context
		attr  string
		index string
	}
	tests := []struct {
		name         string
		mockBehavior func(cl *storagemock.MockStorageClient)
		args         args
		wantErr      error
	}{
		{
			name: "Success",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().DeleteAttr(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			args: args{
				ctx:   context.Background(),
				attr:  "test",
				index: "inner",
			},
		},
		{
			name: "badConnection",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().DeleteAttr(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.Unavailable, "bad connection"))
			},
			args: args{
				ctx:   context.Background(),
				attr:  "test2",
				index: "inner2",
			},
			wantErr: status.Error(codes.Unavailable, "bad connection"),
		},
		{
			name: "indexNotFound",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().DeleteAttr(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.NotFound, "index not found"))
			},
			args: args{
				ctx:   context.Background(),
				attr:  "test2",
				index: "inner2",
			},
			wantErr: status.Error(codes.NotFound, "index not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			cl := storagemock.NewMockStorageClient(c)
			tt.mockBehavior(cl)

			s := &Server{
				tries:   atomic.Uint32{},
				storage: cl,
				ram: RAM{
					available: 100,
					total:     100,
				},
				number: 1,
				mu:     &sync.RWMutex{},
			}
			if err := s.DeleteAttr(tt.args.ctx, tt.args.attr, tt.args.index); (err != nil) && (!errors.Is(err, tt.wantErr)) {
				t.Errorf("DeleteAttr() error = %v, wantCode %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_DeleteIndex(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name         string
		mockBehavior func(cl *storagemock.MockStorageClient)
		args         args
		wantErr      error
	}{
		{
			name: "Success",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().DeleteIndex(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			args: args{
				ctx:  context.Background(),
				name: "test",
			},
		},
		{
			name: "badConnection",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().DeleteIndex(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.Unavailable, "bad connection"))
			},
			args: args{
				ctx:  context.Background(),
				name: "test2",
			},
			wantErr: status.Error(codes.Unavailable, "bad connection"),
		},
		{
			name: "indexNotFound",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().DeleteIndex(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.NotFound, "index not found"))
			},
			args: args{
				ctx:  context.Background(),
				name: "test2",
			},
			wantErr: status.Error(codes.NotFound, "index not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			cl := storagemock.NewMockStorageClient(c)
			tt.mockBehavior(cl)

			s := &Server{
				tries:   atomic.Uint32{},
				storage: cl,
				ram: RAM{
					available: 100,
					total:     100,
				},
				number: 1,
				mu:     &sync.RWMutex{},
			}
			if err := s.DeleteIndex(tt.args.ctx, tt.args.name); (err != nil) && (!errors.Is(err, tt.wantErr)) {
				t.Errorf("DeleteIndex() error = %v, wantCode %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_Get(t *testing.T) {
	type args struct {
		ctx context.Context
		Key string
	}
	tests := []struct {
		name         string
		mockBehavior func(cl *storagemock.MockStorageClient)
		args         args
		want         *storage.GetResponse
		wantErr      error
	}{
		{
			name: "Success",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&storage.GetResponse{Value: "test"}, nil)
			},
			args: args{
				ctx: context.Background(),
				Key: "test",
			},
			want: &storage.GetResponse{
				Value: "test",
			},
		},
		{
			name: "badConnection",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().Get(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.Unavailable, "bad connection"))
			},
			args: args{
				ctx: context.Background(),
				Key: "test2",
			},
			wantErr: status.Error(codes.Unavailable, "bad connection"),
		},
		{
			name: "NotFound",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().Get(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.NotFound, "not found"))
			},
			args: args{
				ctx: context.Background(),
				Key: "test2",
			},
			wantErr: status.Error(codes.NotFound, "not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			cl := storagemock.NewMockStorageClient(c)
			tt.mockBehavior(cl)

			s := &Server{
				tries:   atomic.Uint32{},
				storage: cl,
				ram: RAM{
					available: 100,
					total:     100,
				},
				number: 1,
				mu:     &sync.RWMutex{},
			}
			got, err := s.Get(tt.args.ctx, tt.args.Key)
			if (err != nil) && (!errors.Is(err, tt.wantErr)) {
				t.Errorf("Get() error = %v, wantCode %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetFromIndex(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
		Key  string
	}
	tests := []struct {
		name         string
		mockBehavior func(cl *storagemock.MockStorageClient)
		args         args
		want         *storage.GetResponse
		wantErr      error
	}{
		{
			name: "success",
			args: args{
				ctx:  context.Background(),
				name: "test",
				Key:  "test",
			},
			want: &storage.GetResponse{
				Value: "test",
			},
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().GetFromIndex(gomock.Any(), gomock.Any()).
					Return(&storage.GetResponse{Value: "test"}, nil)
			},
		},
		{
			name: "badConnection",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().GetFromIndex(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.Unavailable, "bad connection"))
			},
			args: args{
				ctx: context.Background(),
				Key: "test2",
			},
			wantErr: status.Error(codes.Unavailable, "bad connection"),
		},
		{
			name: "NotFound",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().GetFromIndex(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.NotFound, "not found"))
			},
			args: args{
				ctx: context.Background(),
				Key: "test3",
			},
			wantErr: status.Error(codes.NotFound, "not found"),
		},
	}
	c := gomock.NewController(t)
	defer c.Finish()
	cl := storagemock.NewMockStorageClient(c)
	for _, tt := range tests {
		tt.mockBehavior(cl)
		s := &Server{
			tries:   atomic.Uint32{},
			storage: cl,
			ram: RAM{
				available: 100,
				total:     100,
			},
			number: 1,
			mu:     &sync.RWMutex{},
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetFromIndex(tt.args.ctx, tt.args.name, tt.args.Key)
			if (err != nil) && (!errors.Is(err, tt.wantErr)) {
				t.Errorf("GetFromIndex() error = %v, wantCode %v", err, tt.wantErr)
				return
			}

			if got == nil && tt.want != nil {
				t.Errorf("GetFromIndex() got = %v, want %v", got, tt.want)
				return
			} else if got == nil || tt.want == nil {
				return
			}

			if !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("GetFromIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetIndex(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name         string
		mockBehavior func(cl *storagemock.MockStorageClient)
		args         args
		want         *storage.GetIndexResponse
		wantErr      error
	}{
		{
			name: "success",
			args: args{
				ctx:  context.Background(),
				name: "TestServer_GetIndex1",
			},
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().GetIndex(gomock.Any(), gomock.Any()).Return(
					&storage.GetIndexResponse{
						Index: map[string]string{
							"Key_GetIndex1": "test",
							"Key_GetIndex2": "test",
							"Key_GetIndex3": "test",
						},
					}, nil,
				)
			},
			want: &storage.GetIndexResponse{
				Index: map[string]string{
					"Key_GetIndex1": "test",
					"Key_GetIndex2": "test",
					"Key_GetIndex3": "test",
				},
			},
		},
		{
			name: "badConnection",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().GetIndex(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.Unavailable, "bad connection"))
			},
			args: args{
				ctx:  context.Background(),
				name: "TestServer_GetIndex2",
			},
			wantErr: status.Error(codes.Unavailable, "bad connection"),
		},
		{
			name: "NotFound",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().GetIndex(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.NotFound, "not found"))
			},
			args: args{
				ctx:  context.Background(),
				name: "TestServer_GetIndex3",
			},
			wantErr: status.Error(codes.NotFound, "not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			cl := storagemock.NewMockStorageClient(c)
			tt.mockBehavior(cl)
			s := &Server{
				tries:   atomic.Uint32{},
				storage: cl,
				ram: RAM{
					available: 100,
					total:     100,
				},
				number: 1,
				mu:     &sync.RWMutex{},
			}
			got, err := s.GetIndex(tt.args.ctx, tt.args.name)
			if (err != nil) && (!errors.Is(err, tt.wantErr)) {
				t.Errorf("IndexToJSON() error = %v, wantCode %v", err, tt.wantErr)
				return
			}
			if tt.wantErr != nil {
				return
			}

			if got != nil && tt.want != nil {
				if !reflect.DeepEqual(*got, *tt.want) {
					t.Errorf("IndexToJSON() got = %v, want %v", got, tt.want)
				}
			} else {
				t.Errorf("IndexToJSON() got = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestServer_NewIndex(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name         string
		args         args
		mockBehavior func(cl *storagemock.MockStorageClient)
		wantErr      error
	}{
		{
			name: "success",
			args: args{
				ctx:  context.Background(),
				name: "TestServer_NewIndex1",
			},
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().NewIndex(gomock.Any(), gomock.Any()).Return(
					nil, nil)
			},
		},
		{
			name: "badConnection",
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().NewIndex(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.Unavailable, "bad connection"))
			},
			args: args{
				ctx:  context.Background(),
				name: "TestServer_NewIndex2",
			},
			wantErr: status.Error(codes.Unavailable, "bad connection"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			cl := storagemock.NewMockStorageClient(c)
			tt.mockBehavior(cl)

			s := &Server{
				tries:   atomic.Uint32{},
				storage: cl,
				ram: RAM{
					available: 100,
					total:     100,
				},
				number: 1,
				mu:     &sync.RWMutex{},
			}
			if err := s.NewIndex(tt.args.ctx, tt.args.name); (err != nil) && (!errors.Is(err, tt.wantErr)) {
				t.Errorf("NewIndex() error = %v, wantCode %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_Set(t *testing.T) {
	type args struct {
		ctx    context.Context
		Key    string
		Value  string
		unique bool
	}
	tests := []struct {
		name         string
		args         args
		mockBehavior func(cl *storagemock.MockStorageClient)
		wantErr      error
	}{
		{
			name: "success",
			args: args{
				ctx:    context.Background(),
				Key:    "Key_Set",
				Value:  "test",
				unique: false,
			},
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().Set(gomock.Any(), gomock.Any()).Return(
					nil, nil)
			},
		},
		{
			name: "badConnection",
			args: args{
				ctx:    context.Background(),
				Key:    "Key_Set",
				Value:  "test",
				unique: false,
			},
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().Set(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.Unavailable, "bad connection"))
			},
			wantErr: status.Error(codes.Unavailable, "bad connection"),
		},
		{
			name: "AlreadyExists",
			args: args{
				ctx:    context.Background(),
				Key:    "Key_Set",
				Value:  "test",
				unique: true,
			},
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().Set(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.AlreadyExists, "already exists"))
			},
			wantErr: status.Error(codes.AlreadyExists, "already exists"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			cl := storagemock.NewMockStorageClient(c)
			tt.mockBehavior(cl)

			s := &Server{
				tries:   atomic.Uint32{},
				storage: cl,
				ram: RAM{
					available: 100,
					total:     100,
				},
				number: 1,
				mu:     &sync.RWMutex{},
			}
			if err := s.Set(tt.args.ctx, tt.args.Key, tt.args.Value, tt.args.unique); (err != nil) && (!errors.Is(err, tt.wantErr)) {
				t.Errorf("Set() error = %v, wantCode %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_SetToIndex(t *testing.T) {
	type args struct {
		ctx    context.Context
		name   string
		Key    string
		Value  string
		unique bool
	}
	tests := []struct {
		name         string
		mockBehavior func(cl *storagemock.MockStorageClient)
		args         args
		wantErr      error
	}{
		{
			name: "success",
			args: args{
				ctx:    context.Background(),
				name:   "TestServer_SetToIndex",
				Key:    "Key_Set",
				Value:  "test",
				unique: false,
			},
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().SetToIndex(gomock.Any(), gomock.Any()).Return(
					nil, nil)
			},
		},
		{
			name: "badConnection",
			args: args{
				ctx:    context.Background(),
				name:   "TestServer_SetToIndex",
				Key:    "Key_Set",
				Value:  "test",
				unique: false,
			},
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().SetToIndex(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.Unavailable, "bad connection"))
			},
			wantErr: status.Error(codes.Unavailable, "bad connection"),
		},
		{
			name: "AlreadyExists",
			args: args{
				ctx:    context.Background(),
				name:   "TestServer_SetToIndex",
				Key:    "Key_Set",
				Value:  "test",
				unique: true,
			},
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().SetToIndex(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.AlreadyExists, "already exists"))
			},
			wantErr: status.Error(codes.AlreadyExists, "already exists"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			cl := storagemock.NewMockStorageClient(c)
			tt.mockBehavior(cl)

			s := &Server{
				tries:   atomic.Uint32{},
				storage: cl,
				ram: RAM{
					available: 100,
					total:     100,
				},
				number: 1,
				mu:     &sync.RWMutex{},
			}
			if err := s.SetToIndex(tt.args.ctx, tt.args.name, tt.args.Key, tt.args.Value, tt.args.unique); (err != nil) && (!errors.Is(err, tt.wantErr)) {
				t.Errorf("SetToIndex() error = %v, wantCode %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_Size(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name         string
		mockBehavior func(cl *storagemock.MockStorageClient)
		args         args
		want         *storage.IndexSizeResponse
		wantErr      error
	}{
		{
			name: "success",
			args: args{
				ctx:  context.Background(),
				name: "TestServer_Size",
			},
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().Size(gomock.Any(), gomock.Any()).Return(
					&storage.IndexSizeResponse{
						Size: 100,
					}, nil)
			},
			want: &storage.IndexSizeResponse{
				Size: 100,
			},
		},
		{
			name: "badConnection",
			args: args{
				ctx:  context.Background(),
				name: "TestServer_Size",
			},
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().Size(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.Unavailable, "bad connection"))
			},
			wantErr: status.Error(codes.Unavailable, "bad connection"),
		},
		{
			name: "NotFound",
			args: args{
				ctx:  context.Background(),
				name: "TestServer_Size",
			},
			mockBehavior: func(cl *storagemock.MockStorageClient) {
				cl.EXPECT().Size(gomock.Any(), gomock.Any()).Return(
					nil, status.Error(codes.NotFound, "not found"))
			},
			wantErr: status.Error(codes.NotFound, "not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			cl := storagemock.NewMockStorageClient(c)
			tt.mockBehavior(cl)

			s := &Server{
				tries:   atomic.Uint32{},
				storage: cl,
				ram: RAM{
					available: 100,
					total:     100,
				},
				number: 1,
				mu:     &sync.RWMutex{},
			}
			got, err := s.Size(tt.args.ctx, tt.args.name)
			if (err != nil) && (!errors.Is(err, tt.wantErr)) {
				t.Errorf("Size() error = %v, wantCode %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Size() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_setRAM(t *testing.T) {
	type args struct {
		ram *storage.AttachToIndexResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				ram: &storage.AttachToIndexResponse{
					Ram: &storage.Ram{
						Available: 100,
						Total:     100,
					},
				},
			},
		},
		{
			name: "nil",
			args: args{
				ram: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				tries: atomic.Uint32{},
				ram: RAM{
					available: 100,
					total:     100,
				},
				number: 1,
				mu:     &sync.RWMutex{},
			}
			s.setRAM(tt.args.ram)

			if tt.args.ram == nil {
				return
			}
			if s.ram.available != tt.args.ram.Ram.Available {
				t.Errorf("setRAM() = %v, want %v", s.ram.available, tt.args.ram.Ram.Available)
			}
			if s.ram.total != tt.args.ram.Ram.Total {
				t.Errorf("setRAM() = %v, want %v", s.ram.total, tt.args.ram.Ram.Total)
			}
		})
	}
}
