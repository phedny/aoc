(ns day-16-f (:require util))

(def input
  (first (util/read-input)))

(defn hex->bits
  [h]
  (let [b (map #(- (byte %) 48) (Integer/toString (Character/digit h 16) 2))]
    (concat (repeat (- 4 (count b)) 0) (seq b))))

(defn string->bit-stream
  [s]
  (mapcat hex->bits (seq s)))

(defn take-number
  [n coll]
  (cons (reduce #(+ (* 2 %1) %2) 0 (take n coll)) (drop n coll)))

(defn take-literal-value
  ([bits]
   (take-literal-value 0 bits))
  ([acc bits]
   (let [[m & bits] (take-number 1 bits)
         [v & bits] (take-number 4 bits)
         result (+ (* acc 16) v)]
     (if (= m 0)
       (cons result bits)
       (recur result bits)))))

(declare take-packet)

(defn take-all-packets
  [bits]
  (let [[packet & bits] (take-packet bits)]
    (if (empty? bits)
      (list packet)
      (cons packet (take-all-packets bits)))))

(defn take-number-of-packets
  [n bits]
  (if (= n 0)
    (cons () bits)
    (let [[packet & bits] (take-packet bits)
          [sub-packets & bits] (take-number-of-packets (dec n) bits)]
      (cons (cons packet sub-packets) bits))))

(defn take-sub-packets
  [bits]
  (let [[length-type & bits] (take-number 1 bits)]
    (if (= length-type 0)
      (let [[length & bits] (take-number 15 bits)]
        (cons (take-all-packets (take length bits)) (drop length bits)))
      (let [[number & bits] (take-number 11 bits)]
        (take-number-of-packets number bits)))))

(defn take-packet
  [bits]
  (let [[version & bits] (take-number 3 bits)
        [type-id & bits] (take-number 3 bits)]
    (if (= type-id 4)
      (let [[value & bits] (take-literal-value bits)]
        (cons {:version version :type-id type-id :value value} bits))
      (let [[sub-packets & bits] (take-sub-packets bits)]
        (cons {:version version :type-id type-id :sub-packets sub-packets} bits)))))

(defn version-sum
  [{version :version sub-packets :sub-packets}]
  (apply + version (map version-sum (or sub-packets ()))))

(println (-> input string->bit-stream take-packet first version-sum))

(defn packet->clojure
  [packet]
  (case (packet :type-id)
    0 (apply list '+ (map packet->clojure (packet :sub-packets)))
    1 (apply list '* (map packet->clojure (packet :sub-packets)))
    2 (apply list 'min (map packet->clojure (packet :sub-packets)))
    3 (apply list 'max (map packet->clojure (packet :sub-packets)))
    4 (packet :value)
    5 (list 'if (apply list '> (map packet->clojure (packet :sub-packets))) 1 0)
    6 (list 'if (apply list '< (map packet->clojure (packet :sub-packets))) 1 0)
    7 (list 'if (apply list '= (map packet->clojure (packet :sub-packets))) 1 0)))

;(println (-> input string->bit-stream take-packet first packet->clojure))
(println (-> input string->bit-stream take-packet first packet->clojure eval))
