package entity

type (
  Pagination struct {
    Page       int64 `json:"page"`
    PerPage    int64 `json:"perpage"`
    TotalCount int64 `json:"total_count"`
  }

  IPagination interface {
    Offset() uint64
    Limit() uint64
    HasNext() bool
    HasPrev() bool
  }
)

func NewPagination(page, perPage int64) *Pagination {
  return &Pagination{
    Page:    page,
    PerPage: perPage,
  }
}

// Offset オフセット値を取得
func (m *Pagination) Offset() uint64 {
  return uint64((m.Page - 1) * m.PerPage)
}

// Limit リミット値を取得
func (m *Pagination) Limit() uint64 {
  return uint64(m.PerPage)
}

// HasNext 次のページがあるかどうか
func (m *Pagination) HasNext() bool {
  return m.TotalCount > (m.Page * m.PerPage)
}

// HasPrev 前のページがあるかどうか
func (m *Pagination) HasPrev() bool {
  return m.Page > 1
}