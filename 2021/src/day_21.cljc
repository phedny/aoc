(ns day-21)

(defn play-step
  [sum-of-rolls player-position other-position player-score other-score step-count]
  (let [new-position (inc (mod (+ player-position sum-of-rolls -1) 10))
        new-score (+ new-position player-score)]
    [other-position new-position other-score new-score (inc step-count)]))

(defn play-practice-game
  [die & args]
  (let [[_ _ _ new-score :as result] (apply play-step (* 3 (+ 2 die)) args)]
    (if (>= new-score 1000)
      (apply vector (+ 3 die) result)
      (recur (+ 3 die) result))))

;(def player1-start 4)
;(def player2-start 8)
(def player1-start 10)
(def player2-start 2)

(let [[die-rolls _ _ loser-score _ _] (play-practice-game 0 player1-start player2-start 0 0 0)]
  (println (* die-rolls loser-score)))

(def initial-universe
  [player1-start player2-start 0 0 0])

(def roll-distribution
  (frequencies (for [a (range 1 4) b (range 1 4) c (range 1 4)] (+ a b c))))

(def play-step (memoize play-step))

(defn finished?
  [[_ _ & scores]]
  (some #(> % 20) scores))

(defn split-universe
  [universe n]
  (if (finished? universe)
    {universe n}
    (reduce #(assoc %1 (apply play-step (first %2) universe) (* n (second %2))) {} roll-distribution)))

(defn play-dirac-game
  [universes]
  (if (every? finished? (keys universes))
    universes
    (let [value (map (partial apply split-universe) universes)]
      (recur (apply merge-with + value)))))

(def final-universes
  (play-dirac-game {initial-universe 1}))

(def player-1-wins
  (apply + (map second (filter #(= (mod (get (first %) 4) 2) 1) final-universes))))

(def player-2-wins
  (apply + (map second (filter #(= (mod (get (first %) 4) 2) 0) final-universes))))

(println (max player-1-wins player-2-wins))
