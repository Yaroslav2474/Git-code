package main

import (
	"database/sql"
	"fmt"
	"time"

	"math/rand/v2"

	_ "github.com/mattn/go-sqlite3"
)

type Hero struct {
	max_hp  int
	name    string
	hp      int
	armor   int
	damage  int
	weapon  string
	crit    int
	gold    int
	exp     float64
	req_exp float64
	lvl     int
}

type Enemy struct {
	name   string
	hp     int
	armor  int
	damage int
	weapon string
	crit   int
	gold   int
	lvl    int
	exp    int
}

func (hr Hero) info() {
	fmt.Println("------------------------------------------------hero info-------------------------------------------------------------------")
	fmt.Printf("Уровень: %d (%d/%d опыта)\nИмя: %s\nХп: %d/%d\nБроня: %d\nУрон: %d\nОружие: %s\nШанс крит.удара: %d\nЗолото: %d\n\n", hr.lvl, int(hr.exp), int(hr.req_exp), hr.name, hr.hp, hr.max_hp, hr.armor, hr.damage, hr.weapon, hr.crit, hr.gold)
	fmt.Println("----------------------------------------------------------------------------------------------------------------------------")

}

func (hr Enemy) enemy_info() {
	fmt.Printf("Уровень: %d (%d опыта)\nИмя: %s\nХп: %d\nБроня: %d\nУрон: %d\nОружие: %s\nШанс крит.удара: %d\nЗолото: %d\n\n", hr.lvl, hr.exp, hr.name, hr.hp, hr.armor, hr.damage, hr.weapon, hr.crit, hr.gold)

}

func (hr *Hero) Dungeon() {

	enemies_defeated, exp_multiplier, total_exp := 0, 1.0, 0.0

	for {
		enemy := hr.gen_enemy()
		fmt.Printf("%s замечает, что вдали сидит %s\n", hr.name, enemy.name)
		enemy.enemy_info()
		fmt.Print("Напасть? (1-да/0-нет): ")
		dmg := 0
		c := 1

		def_armor := hr.armor

		for {
			fmt.Scan(&c)
			if c == 1 {
				if rand.IntN(100)+1 <= hr.crit {
					dmg = hr.damage * 3
					fmt.Println("КРИТИЧЕСКИЙ УРОН")
				} else {
					dmg = hr.damage
				}

				enemy.armor -= dmg
				if enemy.armor <= 0 {
					enemy.hp += enemy.armor
					enemy.armor = 0
				}
				if enemy.hp <= 0 {
					enemies_defeated += 1
					enemy.hp = 0
					fmt.Print("Победа\n")
					hr.gold += enemy.gold
					hr.armor = def_armor
					total_exp += float64(enemy.exp) * exp_multiplier
					hr.exp += float64(enemy.exp) * exp_multiplier
					exp_multiplier += 0.5
					if float64(hr.exp) >= hr.req_exp {
						hr.exp -= hr.req_exp

						hr.req_exp = 1.6 * hr.req_exp
						hr.lvl += 1
						hr.max_hp = 75 + 25*hr.lvl
						hr.damage = 20 + 5*hr.lvl

					}

					break
				}
				fmt.Printf("\n%s нанес %s %d урона.\nТеперь у %s %d здоровья и %d брони.\n\n", hr.name, enemy.name, dmg, enemy.name, enemy.hp, enemy.armor)
				if rand.IntN(100)+1 <= enemy.crit {
					dmg = hr.damage * 3
					fmt.Println("КРИТИЧЕСКИЙ УРОН")

				} else {
					dmg = enemy.damage
				}
				hr.armor -= dmg
				if hr.armor <= 0 {
					hr.hp += hr.armor
					hr.armor = 0
				}
				fmt.Printf("%s нанес %s %d урона.\nТеперь у %s %d здоровья и %d брони.\n", enemy.name, hr.name, dmg, hr.name, hr.hp, hr.armor)

				if hr.hp <= 0 {
					fmt.Print("Проигрышь\n")
					break
				}
			} else if c == 0 {
				break
			} else {
				fmt.Print("\nНет такой комнады!!!\n")

			}
			fmt.Print("\nПродолжить? (1-да/0-нет): ")
		}
		if c == 0 {
			break
		}
	}
	fmt.Printf("Побеждено %d врагов подряд.\nВы заработали %d опыта.\n", enemies_defeated, int(total_exp))
}

func (hr *Hero) market() {

}

func (hr *Hero) boss_batle() {
	fmt.Print("Босс в пути!")
	fmt.Println()
}

func (hr *Hero) gen_enemy() *Enemy {
	var weapons = []string{
		"Палка",
		"Ржавый меч",
		"Кинжал",
		"Меч",
		"Тяжелый топор",
	}

	lvl := rand.IntN(hr.lvl) + 1

	hp := 50 + lvl*25 + rand.IntN(25)
	armor := 5 + lvl*3 + rand.IntN(5)
	damage := 10 + lvl*5 + rand.IntN(10)
	crit := 5 + lvl*2 + rand.IntN(3)
	gold := lvl*10 + rand.IntN(20)
	exp := lvl*15 + rand.IntN(25)

	weapon := weapons[rand.IntN(len(weapons))]

	enemy := &Enemy{
		name:   "Враг",
		hp:     hp,
		armor:  armor,
		damage: damage,
		weapon: weapon,
		crit:   crit,
		gold:   gold,
		lvl:    lvl,
		exp:    exp,
	}
	return enemy
}

func (hr *Hero) sleep() {
	fmt.Println("------------------------------------------------Sleep-----------------------------------------------------------------------")
	fmt.Print("Сон: ")
	for hr.hp < hr.max_hp {
		hr.hp += 2
		time.Sleep(1 * time.Second)
		if hr.hp > hr.max_hp {
			hr.hp -= 1
		}
		fmt.Printf("(%d/%d) ", hr.hp, hr.max_hp)
	}

	fmt.Println("----------------------------------------------------------------------------------------------------------------------------")
}

func (hr *Hero) menu() string {
	var choi int
	for {
		fmt.Println("------------------------------------------------Menu------------------------------------------------------------------------")

		fmt.Printf("Привет, %s!\n\t1.вывести информацию о герое\n\t2.магазино\n\t3.пойти в данж\n\t4.пойти на босса\n\t5.Слиппи тайм\n\t0.выход\n", hr.name)
		fmt.Println("----------------------------------------------------------------------------------------------------------------------------")
		fmt.Scan(&choi)
		switch choi {
		case 1:
			hr.info()
		case 2:
			hr.market()
		case 3:
			hr.Dungeon()
		case 4:
			hr.boss_batle()
		case 5:
			hr.sleep()
		case 0:
			return "Пока"

		}
	}
}

func main() {

	database, _ := sql.Open("sqlite3", "./db.db")

	var name string

	fmt.Print("Введите имя вашего персонажа: ")

	fmt.Scan(&name)

	herro := &Hero{100, name, 100, 50, 20, "ак-47", 50, 100, 0, 100, 1}

	herro.menu()

}
