package chat

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/router"
)

func Run(r *router.Router) {
	users := r.All("user", "order by created_at", "")
	for _, user := range users {
		friend := NewFriend(user)
		go friend.OnlineOffline()
		intSeconds := rand.Intn(10) + 6
		time.Sleep(time.Second * time.Duration(intSeconds))
	}
}

func turnGreenLightOn(path, seconds string) {
	b, err := exec.Command("python3", "python/online.py", "python/"+path, seconds).CombinedOutput()
	fmt.Println(string(b), err)
}

func firstNameLower(s string) string {
	tokens := strings.Split(s, " ")
	return strings.ToLower(tokens[0])
}
