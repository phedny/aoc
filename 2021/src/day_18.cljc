(ns day-18 (:require util))

(def input
  (util/read-input read-string))

(defn add-left
  [sf-number left]
  (if (number? sf-number)
    (+ sf-number left)
    [(add-left (first sf-number) left) (second sf-number)]))

(defn add-right
  [sf-number right]
  (if (number? sf-number)
    (+ sf-number right)
    [(first sf-number) (add-right (second sf-number) right)]))

(defn sf-explode
  ([sf-number]
   (let [result (sf-explode sf-number 0)]
     (if (map? result)
       (:value result)
       result)))
  ([sf-number depth]
   (cond
     (number? sf-number) sf-number
     :else (let [left (sf-explode (first sf-number) (inc depth))]
             (cond
               ;(and (map? left) (number? (second sf-number))) {:value [(:value left) (+ (or (:right left) 0) (second sf-number))] :left (:left left)}
               (map? left) {:value [(:value left) (add-left (second sf-number) (or (:right left) 0))] :left (:left left)}
               :else (let [right (sf-explode (second sf-number) (inc depth))]
                       (cond
                         (map? right) {:value [(add-right left (or (:left right) 0)) (:value right)] :right (:right right)}
                         (< depth 4) [left right]
                         :else {:value 0 :left left :right right})))))))

(defn sf-split
  [sf-number]
  (cond
    (vector? sf-number) (let [left (sf-split (first sf-number))]
                          (if (= left (first sf-number))
                            [left (sf-split (second sf-number))]
                            [left (second sf-number)]))
    (< sf-number 10) sf-number
    :else (let [half-number (/ (double sf-number) 2)]
            [(long (Math/floor half-number)) (long (Math/ceil half-number))])))

(defn sf-reduce
  [sf-number]
  (let [exploded-number (sf-explode sf-number)]
    (if (= sf-number exploded-number)
      (let [split-number (sf-split sf-number)]
        (if (= sf-number split-number)
          sf-number
          (recur split-number)))
      (recur exploded-number))))

(defn sf-plus
  [a b]
  (sf-reduce [a b]))

(defn sf-magnitude
  [sf-number]
  (if (number? sf-number)
    sf-number
    (+
      (* 3 (sf-magnitude (first sf-number)))
      (* 2 (sf-magnitude (second sf-number))))))

(def result-part-a
  (sf-magnitude (reduce sf-plus input)))

(println result-part-a)

(def result-part-b
  (apply max (for [a input
                   b input
                   :when (not= a b)]
               (sf-magnitude (sf-plus a b)))))

(println result-part-b)
