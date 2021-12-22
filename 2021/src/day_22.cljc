(ns day-22 (:require util))
(def filename "example3")

(def input
  (apply util/read-input #"\n" #"(\.\.|[ ,]?[xyz]=)" (partial = "on") (repeat 6 #(Long/parseLong %))))

(defn expand-coords
  [coords]
  (let [ranges (map (fn [[from to]] `(range ~from ~(inc to))) coords)
        seq-exprs (mapcat #(vector (gensym) %) ranges)]
    (eval (list 'for (apply vector seq-exprs) (apply vector (map first (partition 2 seq-exprs)))))))

(defn find-on
  [cubes-on instr]
  (if (empty? instr)
    cubes-on
    (let [[turn-on & coords] (first instr)]
      (recur (apply (if turn-on conj disj) cubes-on (expand-coords (partition 2 coords))) (rest instr)))))

(defn initialization?
  [[_ & coords]]
  (every? #(and (<= -50 (first %)) (>= 50 (second %))) (partition 2 coords)))

(println (count (find-on #{} (filter initialization? input))))
