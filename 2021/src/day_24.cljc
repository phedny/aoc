(ns day-24)

(def spec [11 3, 14 7, 13 1, -4 6, 11 14, 10 7, -4 9, -12 9, 10 6, -11 4, 12 0, -1 7, 0 12, -11 1])

(defn make-constraint
  [a b d]
  (if (neg-int? d) [b a (unchecked-negate d)] [a b d]))

(defn add-input
  [[r [t :as s]] [k [n1 n2]]]
  (let [z (if (pos-int? n1) s (rest s))
        a (+ (or (second t) 0) n1)]
    (if (< -9 a 9)
      [[(conj r (make-constraint (first t) k a)) z] [(conj r (make-constraint (first t) k a)) (conj z [k n2])]]
      [[r (conj z [k n2])]])))

(defn add-inputs
  [ss kn]
  (mapcat #(add-input % kn) ss))

(def inputs
  (map (comp keyword str char) (range (int \a) (int \o))))

(def results
  (reduce add-inputs [[() ()]] (map vector inputs (partition 2 spec))))

(def constraints
  (some #(if (empty? (second %)) (first %)) results))

(defn max-value
  [k [l h d]]
  (cond (= h k) 9 (= l k) (- 9 d)))

(defn min-value
  [k [l h d]]
  (cond (= l k) 1 (= h k) (inc d)))

(defn compute-result
  [f]
  (apply str (map #(some (partial f %) constraints) inputs)))

(println (compute-result max-value))
(println (compute-result min-value))
