(ns day-15 (:require util [clojure.set :refer (union difference)]))

(def chars->ints
  (partial map #(Character/digit % 10)))

(def input
  (apply vector (util/read-input (comp #(apply vector %) chars->ints seq))))

(defn neighbors
  [row col]
  (filter (complement #(nil? (get-in input %))) [[(dec row) col] [(inc row) col] [row (dec col)] [row (inc col)]]))

(defn extend-sum-cost
  [n cost->cell cell->cost]
  (let [computed-cells (mapcat #(map (fn [cell] [cell (+ n (get-in input cell))]) (apply neighbors %)) (cost->cell n))
        new-cells (filter #(not (contains? cell->cost (first %))) computed-cells)]
    [(reduce
       #(update %1 (second %2) (fn [s] (conj (or s #{}) (first %2))))
       cost->cell
       new-cells)
     (reduce
       #(update %1 (first %2) (fn [_] (second %2)))
       cell->cost
       new-cells)]))

(defn sum-cost*
  [n cost->cell cell->cost]
  (lazy-seq (cons cell->cost (apply sum-cost* (inc n) (extend-sum-cost n cost->cell cell->cost)))))

(def sum-cost
  (sum-cost* 0 {0 #{[0 0]}} {[0 0] 0}))

(def end-cell
  [(dec (count input)) (dec (count (first input)))])

(def costs-with-last
  (first (drop-while #(not (contains? % end-cell)) sum-cost)))

(println (get costs-with-last end-cell))
