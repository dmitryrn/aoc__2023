package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = `
........................................................862...........20.............453...619......58........694...312.................292.
...846................132.49........308..........................=............50.....*..............*........+.....+...............59.......
........../46....140.......*............735......852&..706.....860...............297.459..........998................661..418.883.......+...
....*...............*....727......613..#.....517..........-........*..............*.......................888.......*......*...*.........982
.828.865....395......163......................*......381............312....34...533..............................291.....440.488..370.......
.................+............122.....598*....400....+......................+.=.......451........746*...............................*.......
......864......279........203....*731.....335.....41....23..365.&.......659...817.......*.............89.931......672....661*92....72.......
943.....%..................*....................../..............455.....*...........155................*............*......................
........................966.......823.................%702..881........874...494..........................364+..$.......#587....&....157....
.......701.621......................%........%....93...........*................*.........*.&273.................497............814.........
.........*..+............201.863-.........113...-.+...*363....900.............111......586.......289........898.......306....*..............
...*...894.....&...754...*........#803........609...80...............52....................246..=.............#....../....118..........633..
...321.......642...*...764...644/......980..............$726.........*...306........787...*...................................550...........
.......666........902..............946*..........230...........45.245.......%...481*.......203......433......+.........294...%..............
..........*862..........261...184..........744...*...................................................*........782........&...........984....
....992*........678&.....$.....*..................229......60..........239.........*437...*...681....809......................541#...&......
........812.................704........+......*.......33.......434....*...................840...*..................@300...../...........257.
.37..................-.17.......26..793.....71.486...............*..240........261.............605........*..............953................
...............65.724....*..6....*......446.........#854.318...86................*........+922.........546.762.......................137....
........279....*........382...925............*323........................364.....616.&.........................597.......-419..985*...*.....
..........@.370....................735+...597......@....*966.........................699.874/............446...%............................
..................798.........%..................711.........905........*694.........................700*............................492....
......706............+...%....536.$417................913$......@.....39.......+.......................................-................*...
......*........77.......975..................876..........................-954..443.804....&......36*..............51.369..247..........244.
..787.639......../.................$........./.........49..482.........*............/......952.......365....$.......*......*................
....................................377....$.......603*...........&.....540.......$......................369..215.276.....397.856......-....
...........&.........959...81............370..958..............768..............41..853......................*..................*.473..737..
...802$...560...........*...*....................*......286................/596.....#............+..........490...............411...........
..............817.791...101..727.....404.689....606.......*.718.+...233...........=.......59..202..&...............61..................420..
.....+..730...%.........................*............328@.......90....*..........460......-.......155........913..*...881...................
....928.............285....142......../............................819......755..........................609*.....927.*.....................
..........408..816...*......*.........539..850..21.....&...699.............+......#....952..&844.....88................952..................
......................409...539..674.............=.....337...*....59...........+..698./..............*.....67....930.....................185
......775/....726+.................+..259...141*............471...%..........172.....................185.....*.............670...411....*...
...............................=.......*........246.................447....................554...951........201...723...67.#....*......914..
....................482...../...777.78.507..................$..................918............%...@.../..........&...........688............
.457............40..*....564........$..................29...199..%537.............*....................489.....*........730.................
.......220..239....16........269..........220....#782......................592....857.......................435.343.....%...............913.
.............*.................$...487....$..............183=.........481...-..........118..........419....................286......163.*...
.........../..752.....................*...............#...........933....*.........10.....*........*......................*....735..*....315
.........958.........934.............882............51....512....*....769....$.......$...533.....264............$...6....720.........550....
..408...........366...$......................132.........=.....794........615...........................297...60....*.......................
....$.........6...*........&.....292.....&.................%..........106........889.......@..............@.........468.......$...239.......
.......314.........831.130..48......+.655.....757..976......594...924*......932....*....99..759.2.............750.*..........236.....*251...
..451....*................*......................*....*..........................274....*.................868*.....939......................
....*..477....580..........739...743..346..802#...91...531......739*661..................925..........#................894.55...995.........
.178.........@..........90........*.....*....................................139.....855.............583..#....377........*....../..........
........588.............*...65...261....28..10.............22....697+.812.....*.........&..141.............342.*....710.....................
.......*..............255......*............*.....@155..................@..254......#........*.................516................777.......
......876......364..............453.......745...................216..............187...&..686..............782............936...............
...........579....*...#.......................736................*...@..................9.........797..809....*...........=.....983.....*448
............../.700.152....587..........444...&...375%.744#....21.....140.......=.........576........*..*.....420..............*.....399....
........822...............*............./...................#.....377........657....$.856..*...811.658............111..+.....718............
...689...*..............764....................758..........40...*................632.*...574.....................*...850.........626.......
........112.......................833.........*.......&........456.345..707.....*.....35......96..*.............891.......@218.....#........
................714............................850.....131.................@.334.776.....896..*....878.....729.......846...............641..
............121*.....736....737......742.....&..............429.125..104................*......397................................389.......
....................*.........*...............441....$......*...*...@....841*363...*...103.............*......#.....820....306...$..........
.......772..........851.......580...................750...601....38...............715......./........54.832....226..-.......*............861
.282/....*......264.....-.................63*...........................655...223..........360...........................887................
........130......*....208....................2......976.............241...............969......588...........942.....................542....
............782...208.....+....255........................918...476*......606*464....*.....447....=......674......979........+..............
..............*...........979..........@..........@..........%.....................302........*...../...*..........@.......642........937...
.....200...837....=..109.......525..476...../..564.......#......571..........129........295..623...647...825..........*385......-.....*.....
.....*..........275.#....354....+........492............2...............303............+...........................991.......290...830......
...246....912......................-.........104..940......440*931.595=../..........$.........741*......977....272......778*...../..........
.........+.........@.......507.....803.........*...-.........................559*....328..........966.............%.........516..494.=......
....*842..........464.........*........%750...338........253....386=..656........529......662............/..........757...............868...
.258........................843...................413...*.............*.....513......359.%........588*16.359.........@.........#...%........
..........498.......$..555......@....................*.....*10.....145........%........+......*.............................136.....980.....
.....681......625.147.....*.....588..82%.397........376.677....633.........................523...607.372................................855.
.....*.......*........439..311..............*.....................-.96.......938*815.839...........-.$....................%.........286.....
...830........115......*.......908.........351...850....290..959/...*...................*.....*453................392......750.........*954.
.....................-.627.......................+..................120............50.900..........818.....668......*.............848.......
.....979*.....603.863......250.......*507.*751......411..../350...........@......../.........702...*...292..*........615...........*........
...........................#......676...........140....*..................800.........$..895....#..93..*....51..295.........=....59.........
...............594.801.84................335.....*..521...883.......=33.......801....917....&..........667.....*.........893..........485...
........................+......................55........*....................-................#................895..................*......
.........315.574..168...............................%.215.....123...........................966.........+...........774...526......11.......
...........*...*...-........533...&766...........709............$....931....624......275.............@..189......$....*......#..........=836
..........342..551.........+...........178..621...........561..................&.....*..............214.......207...822.....................
.................................63...+......=......615..*.............499*825....167...................870.....................265.891.....
......................758..........%.....+................953......&...................727/.....273............413...............=.....*818.
....275..........81.....*.447.506/...67..242............+.........793...531*885...18...............*710........*...#....906.................
.......-..645.....-.........*......#.........98*520.565..958.....................*............88..............66..911.............44*.......
.........%.................455...@..298.......................=....510......-...366..162......*..........306.............354.........153....
..................@....53......689...........#.....690.......318..*.........701............156....%519..*....386..917..........#............
......718..754....606...............164.....28.303....*...........921..............342.558..................*......$..67.......682..........
..............*.........586.....541*...........#.....756..931.....................*.......*.....92.......826..........-.....................
662...........775.........*................../.............*...................419........510..$....................*........-822......-....
...*..........................690.398....98...9...999....617............681.........360.............&......*508..594.437..............385...
..647...709...............71....*...*.............$......................*..........*...539*864......826..................@.................
.......$......529...........$.260..672.538+........................996.560.........639....................867.......922....891........695...
..........$..*................................376..............922*............................875...........*.......*...........%....*.....
....331.542..........932..409$..........815..*......................888.322..............205.....=.940..2#.......................411.635....
....................*..........698*119.......65.511................*....+....-......34......=........*......$..574...971.85.*630............
...104.311.....561#..556...........................*514.............663.....416.......................243.728.....&.....*...................
...........521...........*131...979....227..556...........499....-..............895...42.........%..............=.........../............898
............*..........*........../...*.....*........=....*.......614.....498..*......*...........612.205...138.664........190..............
.........834...348..887.730...........475.404.....474..168..620.............$..175...549...............*.....*...................11..459....
.....320........-................401..........261....................................................267......676..=629....-487.............
......./..%142...............................*....................551................-......961.751.........&......................360......
.............................41&.504......416.............144........+..........&.....438.......*...385....635.330................*.........
....544.369..349..&.....=479..........995......566.80......*....595-....=..468..507............301..&.............*619...383.....498........
.......*......*....950................*..........*..=....484.........258...-..............926..............=................*...............
...............224..............*394...81.....904..............................*.........*................537......521...882...231.....*837.
.......370...........625.....116................../11.866..288.........311.....375.708..67...*......380........170../...........&...........
...600*..........811..*....=...........287................................*211......=.......987....*.......448......................-....92.
............408.....+.405...17....32...+.................180...81................*.....*.........315......*..........-..............148.....
.......*401...-....................&.....713.67......475*......*....707.852*...830..484.77.............902........505..=273.................
..837...........................33........*...............492/.80...&......................203...686..........722...............178.........
...*......191.791..........433....$....342..................................689........988..*....#.........@.....#.923*622......*...........
.850.679.%....*.......376.-.....................308.812...742.........-.148....@.........$.663.........818..177..................839...388..
......./...269...496...*........-....274*............*.....*........791..*........@....%.......*190.............*..241....127.........*.....
.74.................*........912.........435..107..790...640...200......578.112..83....252..269...........970.823.....................875...
...%..787.........515.305..........274.&.........*............*.......................................261...*..............+........&.......
...........136.-........*.564.......*..635.....26../....438.469.................@............193.......*.....55.....244#...185..258.287.....
............*...556...............245......855....928....*.................*61.370..74......-.........488...................................
.....970....315...........999.65..........=...........928.....164.......431.........*...........370.....................#...909...468...471.
.961..@..................%......../350.............*................746.............38....261......*.......301...278..750.....*......*.*....
....*.................7.....................%...860.900....300.........*410..............*.........209.+98..*...............801....313.491..
...845......%160.......=......=.....*......313.............*............................547...986............139.331../401..................
........513........362...540..913..655.510..............483......@...#..670..651..433/...........*264.....43.....*............176...........
211.892....*..253.....%...*...............*.....................645..16..................116...............%...309......*569....*.......+...
.........402....*.83$...884......618.......465.............989..........................=....505..%..................101......974....871....
...4..........565...................*....&.......297&.......*...........@....179..637...........*.199...924...-..713........&.....@.........
...................@.367&.#308.....752...997.23...........396..........883....&..............306...........*.148../.........76.979...382-...
.................677..........................*..933..................................637...............659...........678...................
.....*...............445..........952...837.....*............755........795.....#.....=.....178*731..............855.....=..................
......173.15.417......./.....180....*....*....328...607.........*.392...-.....495.546...............183...&.....-..................224......
..........*.....*.............*......16...65.................772...*..............+........937...........981......92.....+.....164......833.
.116.....149.225....*....287....................136..............240..242.749.............*....688.................=......49...-.......*....
...&.............202.985..-...732*812.......................-../.....=......-.798.......533........468...30%.........................228....
.........94...................................+798..&.....431...425............................929*............+............................
.........*.......................704...982...........332.................447............307%................704...775%.......$973...90..424.
..........711..253.....#...........*......*119..............486........#...........89...........438..666...............487..........*.......
.....947..........*.....823..544....883............291.....@..........813............*.425..659*..........28..491.....+..........896........
.............../...943......................590...*....145....414*969..............896....*.........622.........*.938.........=.............
....@...#184.935.............................*.....71......$................84*.........825.....637...*.......528...*......990..............
..745...............534..............58....901...........974...................637.....................87..........361......................
`

func main() {
	matrix := inputToMatrix(input)

	sum := 0
	gearCoordsToNums := map[[2]int][]int{}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if isNumber(matrix[y][x]) {
				ln := 0
				nstr := make([]rune, 1, 3) // could be on stack if used an array
				nstr[0] = matrix[y][x]
				for i := x + 1; inBounds(len(matrix[0]), len(matrix), x, i) && isNumber(matrix[y][i]); i++ {
					ln = i - x
					nstr = append(nstr, matrix[y][i])
				}

				gearCoords, isAdjacent := isAdjacentToSymbol(matrix, x, y, ln)
				if isAdjacent {
					n, err := strconv.Atoi(string(nstr))
					if err != nil {
						panic("insane")
					}
					for _, coord := range gearCoords {
						gearCoordsToNums[coord] = append(gearCoordsToNums[coord], n)
					}
					sum += n
				}

				x += ln + 1
			}
		}
	}

	fmt.Println(sum)

	gearRatioSum := 0
	for _, ns := range gearCoordsToNums {
		if len(ns) != 2 {
			continue
		}

		gearRatioSum += ns[0] * ns[1]
	}

	println(gearRatioSum)
}

func inputToMatrix(s string) [][]rune {
	trimmed := strings.TrimSpace(s)

	lines := strings.Split(trimmed, "\n")

	matrix := make([][]rune, len(lines))

	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	return matrix
}

func isAdjacentToSymbol(matrix [][]rune, startX, startY, length int) ([][2]int, bool) {
	gearCoords := [][2]int{}
	isAdjacent := false

	for x := startX - 1; x <= startX+length+1; x++ {
		for y := startY - 1; y < startY+2; y++ {
			if !inBounds(len(matrix[0]), len(matrix), x, y) {
				continue
			}

			if matrix[y][x] == 42 {
				gearCoords = append(gearCoords, [2]int{x, y})
			}

			if !isNumber(matrix[y][x]) && matrix[y][x] != 46 {
				isAdjacent = true
			}
		}
	}

	return gearCoords, isAdjacent
}

func inBounds(width, height, x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x > width-1 || y > height-1 {
		return false
	}
	return true
}

func isNumber(c rune) bool {
	return c >= 48 && c < 58
}
