(ns day-17)

;(def input {:x [20 30] :y [-10 -5]})
(def input {:x [88 125] :y [-157 -103]})

(defn overshot?
  [x y]
  (or (> x (second (input :x))) (< y (first (input :y)))))

(defn in-target?
  [x y]
  (and (>= x (first (input :x))) (<= x (second (input :x))) (>= y (first (input :y))) (<= y (second (input :y)))))

(defn to-zero
  [x]
  (cond
    (> x 0) (dec x)
    (< x 0) (inc x)
    :else x))

(defn trajectory
  ([vx vy]
   (into [] (trajectory 0 0 vx vy)))
  ([x y vx vy]
   (cond
     (in-target? x y) (list [x y])
     (overshot? x y) ()
     :else (lazy-seq (cons [x y] (trajectory (+ x vx) (+ y vy) (to-zero vx) (dec vy)))))))

(def min-x
  (-> input :x first (* 2) Math/sqrt int))

(def max-x
  (second (input :x)))

(def min-y
  (first (input :y)))

(def max-y
  (unchecked-negate min-y))

(def trajectories
  (for [x (range min-x (inc max-x))
        y (range min-y (inc max-y))
        :let [t (trajectory x y)]
        :when (apply in-target? (last t))]
    t))

(def result-part-a
  (apply max (mapcat #(map second %) trajectories)))

(println result-part-a)
