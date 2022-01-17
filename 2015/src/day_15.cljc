(ns day-15)

(def sprinkles [2 0 -2 0 3])
(def butterscotch [0 5 -3 0 3])
(def chocolate [0 0 5 -1 8])
(def candy [0 -1 0 5 8])

(def candidates
  (for [n_sprinkles (range 0 101)
        :let [p_sprinkles (map #(* % n_sprinkles) sprinkles)]
        n_butterscotch (range 0 (- 101 n_sprinkles))
        :let [p_butterscotch (map #(* % n_butterscotch) butterscotch)]
        n_chocolate (range 0 (- 101 n_sprinkles n_butterscotch))
        :let [p_chocolate (map #(* % n_chocolate) chocolate)
              n_candy (- 100 n_sprinkles n_butterscotch n_chocolate)
              p_candy (map #(* % n_candy) candy)]]
    (map (comp (partial max 0) +) p_sprinkles p_butterscotch p_chocolate p_candy)))

(defn pick-best
  [candidates]
  (apply max (map (comp (partial apply *) drop-last) candidates)))

(println (pick-best candidates))
(println (pick-best (filter #(= 500 (last %)) candidates)))
