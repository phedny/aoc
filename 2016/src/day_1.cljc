(ns day-1 (:require util))

(def input
  (util/read-input #", "
                   (juxt
                     (comp keyword str first)
                     (comp parse-long #(subs % 1)))))

(defn manhattan-distance
  [[x1 y1] [x2 y2]]
  (+ (abs (- x1 x2)) (abs (- y1 y2))))

(defn step-a
  [[x y dx dy] [t n]]
  (let [[dx dy] (case t
                  :L [(unchecked-negate dy) dx]
                  :R [dy (unchecked-negate dx)])]
    [(+ x (* dx n)) (+ y (* dy n)) dx dy]))

(println (manhattan-distance [0 0] (subvec (reduce step-a [0 0 0 1] input) 0 2)))

(defn step-b
  [[x y dx dy seen dups] i]
  (case i
    :L [x y (unchecked-negate dy) dx seen dups]
    :R [x y dy (unchecked-negate dx) seen dups]
    :S [(+ x dx) (+ y dy) dx dy (conj seen [x y]) (if (contains? seen [x y]) (conj dups [x y]) dups)]))

(def steps
  (mapcat (fn [[i n]] (conj (repeat n :S) i)) input))

(println (manhattan-distance [0 0] (last (nth (reduce step-b [0 0 0 1 #{} ()] steps) 5))))
