(ns day-15 (:require util [clojure.set :refer (union difference)]))

(def chars->ints
  (partial map #(Character/digit % 10)))

(def input
  (apply vector (util/read-input (comp #(apply vector %) chars->ints seq))))

(defn neighbors
  [grid row col]
  (filter (complement #(nil? (get-in grid %))) [[(dec row) col] [(inc row) col] [row (dec col)] [row (inc col)]]))

(defn extend-sum-cost
  [grid n cost->cell cell->cost]
  (let [computed-cells (mapcat #(map (fn [cell] [cell (+ n (get-in grid cell))]) (apply neighbors grid %)) (cost->cell n))
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
  [grid n cost->cell cell->cost]
  (lazy-seq (cons cell->cost (apply sum-cost* grid (inc n) (extend-sum-cost grid n cost->cell cell->cost)))))

(defn find-cheapest-path
  [grid]
  (let [sum-cost (sum-cost* grid 0 {0 #{[0 0]}} {[0 0] 0})
        end-cell [(dec (count grid)) (dec (count (first grid)))]
        costs-with-last (first (drop-while #(not (contains? % end-cell)) sum-cost))]
    (get costs-with-last end-cell)))

(println (find-cheapest-path input))

(defn extend-grid-line
  [grid-line increment]
  (let [width (count grid-line)]
    (apply vector
           (map
             #(inc (mod (+ (dec (get grid-line (mod % width))) (quot % width) increment) 9))
             (range (* 5 (count grid-line)))))))

(defn extend-grid
  [grid]
  (let [height (count grid)]
    (apply vector
           (map
             #(extend-grid-line (get grid (mod % height)) (quot % height))
             (range (* 5 (count grid)))))))

(println (find-cheapest-path (extend-grid input)))
