(ns day-21)

(defn play-game
  [die player-position other-position player-score other-score]
  (let [sum-of-rolls (* 3 (+ 2 die))
        new-position (inc (mod (+ player-position sum-of-rolls -1) 10))
        new-score (+ new-position player-score)]
    (if (>= new-score 1000)
      [(+ 3 die) other-position new-position other-score new-score]
      (recur (+ 3 die) other-position new-position other-score new-score))))

;(def player1-start 4)
;(def player2-start 8)
(def player1-start 10)
(def player2-start 2)

(let [[die-rolls _ _ loser-score _] (play-game 0 player1-start player2-start 0 0)]
  (println (* die-rolls loser-score)))
