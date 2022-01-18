(ns day-18 (:require util))

(def input
  (util/read-input (comp (partial map (partial = \#)) seq)))

(defn neighbor-count
  [grid]
  (map
    (fn
      [line & neighbor-lines]
      (apply map
             (comp count (partial filter identity) vector)
             (conj
               (mapcat
                 (fn [line] [line (concat (rest line) [false]) (concat [false] line)])
                 neighbor-lines)
               (concat (rest line) [false])
               (concat [false] line))))
    grid
    (concat (rest grid) [(repeat false)])
    (concat [(repeat false)] grid)))

(defn apply-mask
  [mask grid]
  (map (partial map #(or %1 %2)) grid mask))

(defn grid->mask
  [grid]
  (let [with-corners [(concat [true] (repeat (- (count (first grid)) 2) false) [true])]]
    (concat with-corners (repeat (- (count grid) 2) (repeat (count (first grid)) false)) with-corners)))

(defn step
  [mask grid]
  (map (partial map #(or (= %2 3) (and %1 (= %2 2)) %3)) grid (neighbor-count grid) mask))

(defn on-after
  [input mask steps]
  (let [final-state (nth (iterate (partial step mask) (apply-mask mask input)) steps)]
    (apply + (map #(count (filter identity %)) final-state))))

(println (on-after input (repeat (repeat false)) 100))
(println (on-after input (grid->mask input) 100))
