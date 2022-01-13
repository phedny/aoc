(ns day-25 (:require util))

(def input (vec (util/read-input)))
(def height (count input))
(def width (count (first input)))

(def cucumbers
  (apply merge (for [row (range height)
                     col (range width)
                     :let [val (get-in input [row col])]
                     :when (not= val \.)]
                 {[row col] (if (= val \v) :south :east)})))

(defmulti move (fn [[_ t]] t))
(defmethod move :east [[[row col]]] [[row (mod (inc col) width)] :east])
(defmethod move :south [[[row col]]] [[(mod (inc row) height) col] :south])

(defn maybe-move
  [cs [_ t :as c] d]
  (if (= t d)
    (let [[np _ :as nc] (move c)]
      (if (contains? cs np) c nc))
    c))

(defn step
  [cs]
  (letfn [(do-move [cs d] (into {} (map #(maybe-move cs % d) cs)))]
    (do-move (do-move cs :east) :south)))

(defn steps-to-stable
  [cs n]
  (let [ncs (step cs)]
    (if (= cs ncs)
      n
      (recur ncs (inc n)))))

(println (steps-to-stable cucumbers 1))
