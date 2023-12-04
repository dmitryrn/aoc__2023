use std::collections::HashSet;

fn main() {
    let input = "
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
";

    let lines = input.trim().lines();
    let mut card_to_count: Vec<usize> = vec![1; lines.clone().count()];

    let mut i: isize = -1;
    for line in lines {
        i += 1;

        let card = parse_card(line);

        for j in (i + 1)..(i + 1 + card.match_count as isize) {
            let copies_of_this_card = card_to_count[i as usize] - 1;
            card_to_count[j as usize] += 1 + copies_of_this_card;
        }
    }

    let total: usize = card_to_count.iter().sum();
    println!("{}", total)
}

#[derive(Debug, Copy, Clone)]
struct Card {
    match_count: usize,
}

fn parse_card(card_str: &str) -> Card {
    let relevant = card_str.split(": ").last().unwrap();
    let mut spl = relevant.split(" | ");
    let winning_nums = spl.next().unwrap();
    let nums = spl.last().unwrap();

    let mut winning_nums_map: HashSet<&str> = HashSet::new();
    for num_str in winning_nums.split(' ') {
        if num_str == "" {
            continue;
        }
        winning_nums_map.insert(num_str);
    }

    let mut matches = 0;

    for num_str in nums.split(' ') {
        if num_str == "" {
            continue;
        }
        if winning_nums_map.contains(num_str) {
            matches += 1;
        }
    }

    return Card {
        match_count: matches,
    };
}

#[test]
fn test_parse_card() {
    let c = parse_card("Card   1: 66 92  4 54 39 76 49 27 61 56 | 66 59 85 54 61 86 37 49  6 18 81 39  4 56  2 48 76 72 71 25 27 67 10 92 13");
    assert_eq!(c.match_count, 10);
    let c = parse_card("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53");
    assert_eq!(c.match_count, 4);
}
