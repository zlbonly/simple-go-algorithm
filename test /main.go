package main

import "fmt"

/*type server int

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("hello world!"))
}

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if (r - l + 1 < len) {
				len = r - l + 1
				ansL, ansR = l, l + len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}


	fmt.Sprint(cnt)

	return s[ansL:ansR]
}





func test_gcflags(){
	s := "hello"
	fmt.Println(s)
}*/

func Example(sclie []string, str string, i int) {
	panic("want stack trace")
}

type User struct {
	Username string
}

func Call1(user *User) {
	fmt.Println("%v", user)
}

func main() {
	a := "aaaaa"
	u := &User{Username: a}
	Call1(u)
}
