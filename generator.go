package korwin

import (
	"math/rand"
	"strings"
	"time"
)

var korwin1 = []string{
	"Ja chcę powiedzieć jedną rzecz:",
	"Trzeba powiedzieć jasno:",
	"Jak powiedział wybitny krakowianin Stanisław Lem,",
	"Proszę mnie dobrze zrozumieć:",
	"Ja chciałem państwu przypomnieć, że",
	"Niech państwo nie mają złudzeń:",
	"Powiedzmy to wyraźnie:",
}

var korwin2 = []string{
	"przedstawiciele czerwonej hołoty",
	"ci wszyscy (tfu!) geje",
	"funkcjonariusze reżymowej telewizji",
	"tak zwani ekolodzy",
	"ci wszyscy (tfu!) demokraci",
	"agenci bezpieki",
	"feminazistki",
}

var korwin3 = []string{
	"zupełnie bezkarnie",
	"całkowicie bezczelnie",
	"o poglądach na lewo od komunizmu",
	"celowo i świadomie",
	"z premedytacją",
	"od czasów Okrągłego Stołu",
	"w ramach postępu",
}

var korwin4 = []string{
	"nawołują do podniesienia podatków",
	"próbują wyrzucić kierowców z miast",
	"próbują skłócić Polskę z Rosją",
	"głoszą brednie o globalnym ociepleniu",
	"zakazują posiadania broni",
	"nie dopuszczają prawicy do władzy",
	"uczą dzieci homoseksualizmu",
}

var korwin5 = []string{
	"bo dzięki temu moga kraść",
	"bo dostają za to pieniądze",
	"bo tak się uczy w państwowej szkole",
	"bo bez tego (tfu!) demokracja nie może istnieć",
	"bo głupich jest więcej niż mądrych",
	"bo chcą stworzyć raj na ziemi",
	"bo chcą niszczyć cywilizacje białego człowieka",
}

var korwin6 = []string{
	"przez kolejne kadencje",
	"o czym się nie mówi",
	"i właśnie dlatego Europa umiera",
	"ale przyjdą muzułmanie i zrobią porządek",
	"- tak samo zresztą jak za Hitlera",
	"- proszę zobaczyć, co się dzieje na Zachodzie, jeśli mi państwo nie wierzą",
	"co lat temu sto nikomu nie przyszłoby nawet do głowy",
}

func GenerateStatement() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	statement := []string{
		korwin1[random.Intn(len(korwin1))],
		korwin2[random.Intn(len(korwin2))],
		korwin3[random.Intn(len(korwin3))],
		korwin4[random.Intn(len(korwin4))],
		korwin5[random.Intn(len(korwin5))],
		korwin6[random.Intn(len(korwin6))],
	}
	return strings.Join(statement, " ")
}
