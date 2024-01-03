package token

import (
	"github.com/JPratama7/util"
	"time"
)

type Generator[ID ~string, T util.IdGetter[ID]] struct {
	public, private string
	duration        time.Duration
}

func NewGenerator[ID ~string, T util.IdGetter[ID]](public, private string, duration time.Duration) *Generator[ID, T] {
	return &Generator[ID, T]{
		public:   public,
		private:  private,
		duration: duration,
	}
}
