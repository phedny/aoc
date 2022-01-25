(ns day-21)

(defn who-wins?
  ([bhp bd ba php pd pa] (who-wins? bhp php (max 1 (- pd ba)) (max 1 (- bd pa))))
  ([bhp php bhpd phpd]   (cond
                           (<= bhp 0) :player
                           (<= php 0) :boss
                           :else (recur (- bhp bhpd) (- php phpd) bhpd phpd))))

(def weapons [[8 4 0] [10 5 0] [25 6 0] [40 7 0] [74 8 0]])
(def armors [[0 0 0 :no-armor] [13 0 1] [31 0 2] [53 0 3] [75 0 4] [102 0 5]])
(def rings [[0 0 0 :no-ring-1] [0 0 0 :no-ring-2] [25 1 0] [50 2 0] [100 3 0] [20 0 1] [40 0 2] [80 0 3]])

(def candidates
  (for [weapon weapons
        armor armors
        ring1 rings
        ring2 rings
        :when (not= ring1 ring2)
        :let [items (map + weapon armor ring1 ring2)]]
    [(who-wins? 104 8 1 100 (nth items 1) (nth items 2)) items weapon armor ring1 ring2]))

(println (apply min (map #(first (second %)) (filter #(= (first %) :player) candidates))))
(println (apply max (map #(first (second %)) (filter #(= (first %) :boss) candidates))))
