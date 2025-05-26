package main

import (
	"fmt"

	"math/rand/v2"
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
	fmt.Printf("Уровень: %d (%d/%d опыта)\nИмя: %s\nХп: %d/%d\nБроня: %d\nУрон: %d\nОружие: %s\nШанс крит.удара: %d\nЗолото: %d\n\n", hr.lvl, hr.exp, hr.req_exp, hr.name, hr.hp, hr.max_hp, hr.armor, hr.damage, hr.weapon, hr.crit, hr.gold)

}

func (hr Enemy) enemy_info() {
	fmt.Printf("Имя: %s\nХп: %d\nБроня: %d\nУрон: %d\nОружие: %s\nШанс крит.удара: %d\nЗолото: %d\n\n", hr.name, hr.hp, hr.armor, hr.damage, hr.weapon, hr.crit, hr.gold)

}

func (hr *Hero) Dungeon(enemy *Enemy) {
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
				enemy.hp = 0
				fmt.Print("Победа\n")
				hr.max_hp += rand.IntN(11) + 5
				hr.damage += rand.IntN(5) + 1
				hr.gold += enemy.gold
				hr.armor = def_armor

				if float64(hr.exp) >= hr.req_exp {
					hr.exp -= hr.req_exp

					hr.req_exp = 1.6 * hr.req_exp
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
			if hr.hp <= 0 {
				fmt.Print("Проигрышь\n")
				break
			}
			fmt.Printf("%s нанес %s %d урона.\nТеперь у %s %d здоровья и %d брони.\n", enemy.name, hr.name, dmg, hr.name, hr.hp, hr.armor)

		} else if c == 0 {
			break
		} else {
			fmt.Print("\nНет такой комнады!!!\n")

		}
		fmt.Print("\nПродолжить? (1-да/0-нет): ")
	}

}

func (hr *Hero) market_armor() {
	fmt.Print("Броню заменяют!")
	fmt.Println()
}

func (hr *Hero) market_gun() {
	fmt.Print("Оружие еще не привезли!")
	fmt.Println()
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

	// lvl = rand.IntN(hr.lvl)+1
	// enemy := Enemy{"хохол", rand.IntN(76) + 50, rand.IntN(31) + 10, rand.IntN(21) + 10, "палка", rand.IntN(25), rand.IntN(66) + 5, lvl}
	// return &enemy
	lvl := rand.IntN(hr.lvl) + 1 // уровень врага от 1 до hr.lvl

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

func (hr *Hero) menu() string {
	var choi int
	for {
		fmt.Printf("Привет, %s!\n\t1.вывести информацию о герое\n\t2.купить новое оружие за золото\n\t3.купить новую броню за золото\n\t4.пойти в данж\n\t5.пойти на босса\n\t0.выход\n", hr.name)
		fmt.Scan(&choi)
		switch choi {
		case 1:
			hr.info()
		case 2:
			hr.market_gun()
		case 3:
			hr.market_armor()
		case 4:
			hr.Dungeon(hr.gen_enemy())
		case 5:
			hr.boss_batle()
		case 0:
			return "Пока"
		}
	}
}

func main() {

	herro := &Hero{100, "Ярик", 100, 50, 20, "ак-47", 50, 100, 0, 100, 1}

	herro.menu()

}
