(ns day-22 (:require util))

(def input
  (apply util/read-input #"\n" #"(\.\.|[ ,]?[xyz]=)" (partial = "on") (repeat 6 #(Long/parseLong %))))

(defn ranges-overlap
  [[from1 to1] [from2 to2]]
  (cond
    (< from2 from1) (recur [from2 to2] [from1 to1])
    (<= from2 to1) [from2 (Math/min to1 to2)]))

(defn cubes-overlap
  [cube1 cube2]
  (let [overlapping-ranges (map ranges-overlap cube1 cube2)]
    (if (some nil? overlapping-ranges) nil overlapping-ranges)))

(defn split-range
  [[r-from r-to] [s-from s-to :as split]]
  [[r-from (dec s-from)] split [(inc s-to) r-to]])

(defn merge-cubes
  ([cube tryable-cubes unmergeable-cubes]
   (if (empty? tryable-cubes)
     [cube unmergeable-cubes]
     (let [trying-cube (first tryable-cubes)
           matching-ranges (map = cube trying-cube)
           connecting-ranges (map #(or (= (inc (second %1)) (first %2)) (= (inc (second %2)) (first %1))) cube trying-cube)]
       (if (and (= (dec (count cube)) (count (filter identity matching-ranges))) (= 1 (count (filter identity connecting-ranges))))
         (recur (map (fn [[from1 to1] [from2 to2]] [(Math/min from1 from2) (Math/max to1 to2)]) cube trying-cube) (rest tryable-cubes) unmergeable-cubes)
         (recur cube (rest tryable-cubes) (conj unmergeable-cubes trying-cube))))))
  ([tryable-cubes merged-cubes]
   (if (empty? tryable-cubes)
     merged-cubes
     (let [[merged-cube unmergeable-cubes] (merge-cubes (first tryable-cubes) (rest tryable-cubes) [])]
       (recur unmergeable-cubes (cons merged-cube merged-cubes)))))
  ([cubes]
   (merge-cubes cubes ())))

(defn split-for-subcube
  [cube subcube]
  (let [[x-ranges y-ranges z-ranges] (map split-range cube subcube)
        all-subcubes (for [x x-ranges y y-ranges z z-ranges] [x y z])
        valid-subcubes (filter #(not-any? (fn [[from to]] (> from to)) %) all-subcubes)
        subcubes-without-input (filter (partial not= subcube) valid-subcubes)]
    (cons subcube (merge-cubes subcubes-without-input))))

(defn process-instructions
  [cubes-on instructions]
  (if (empty? instructions)
    cubes-on
    (let [instruction (first instructions)
          turn-on? (first instruction)
          instruction-cube (partition 2 (rest instruction))]
      (if (contains? cubes-on instruction-cube)
        (recur (if turn-on? cubes-on (disj cubes-on instruction-cube)) (rest instructions))
        (if-let [existing-cube (some #(if (cubes-overlap instruction-cube %) %) cubes-on)]
          (let [overlapping-cube (cubes-overlap instruction-cube existing-cube)
                [_ & split-existing-cube] (split-for-subcube existing-cube overlapping-cube)
                existing-cube-overlap-removed (apply conj (disj cubes-on existing-cube) split-existing-cube)]
            (recur existing-cube-overlap-removed instructions))
          (recur (if turn-on? (conj cubes-on instruction-cube) cubes-on) (rest instructions)))))))

(defn count-on
  [cubes-on]
  (apply + (map (fn [[& ranges]] (apply * (map #(- (second %) (first %) -1) ranges))) cubes-on)))

(defn initialization?
  [[_ & coords]]
  (every? #(and (<= -50 (first %)) (>= 50 (second %))) (partition 2 coords)))

(println (count-on (process-instructions #{} (filter initialization? input))))
(println (count-on (process-instructions #{} input)))
