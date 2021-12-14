(ns day-14 (:require util))

(def input
  (util/read-input))

(def initial-polymer
  (-> input first seq))

(def pair-insertion-map
  (let [insertion-pairs (map #(clojure.string/split % #" -> ") (drop 2 input))]
    (reduce
      (fn
        [m [from to]]
        (assoc m (apply vector (seq from)) (first to)))
      {}
      insertion-pairs)))

(defn split-pair
  [[a c]]
  (let [b (pair-insertion-map [a c])]
    [[a b] [b c]]))

(defn expand-pair-frequencies
  [frequencies]
  (let [raw-frequencies (mapcat #(map (fn [p] [p (second %)]) (split-pair (first %))) frequencies)]
    (reduce #(assoc %1 (first %2) (apply + (map second (second %2)))) {} (group-by first raw-frequencies))))

(defn pair-frequencies*
  [frequencies]
  (lazy-seq (cons frequencies (pair-frequencies* (expand-pair-frequencies frequencies)))))

(def pair-frequencies
  (pair-frequencies* (frequencies (map vector initial-polymer (rest initial-polymer)))))

(defn pair-frequencies->letter-frequencies
  [pair-frequencies]
  (reduce
    (fn
      [lc [[_ l] c]]
      (assoc lc l (+ (get lc l 0) c)))
    {(first initial-polymer) 1}
    pair-frequencies))

(def letter-frequencies
  (map pair-frequencies->letter-frequencies pair-frequencies))

(defn max-minus-min
  [m]
  (apply - (map (comp second #(apply % second m)) [max-key min-key])))

(println (max-minus-min (nth letter-frequencies 10)))
(println (max-minus-min (nth letter-frequencies 40)))
