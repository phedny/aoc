(ns day-16 (:require util))

(def input
  (first (util/read-input)))

(defn hex->bits
  [h]
  (let [b (map #(- (byte %) 48) (Integer/toString (Character/digit h 16) 2))]
    (concat (repeat (- 4 (count b)) 0) (seq b))))

(defn string->byte-array
  [s]
  (byte-array (* 4 (count s)) (mapcat hex->bits (seq s))))

(defn string->byte-buffer
  [s]
  (java.nio.ByteBuffer/wrap (string->byte-array s)))

(defn read-number
  [bits bb]
  (reduce (fn [acc _] (+ (* acc 2) (.get bb))) 0 (range bits)))

(defn read-literal
  ([bb] (read-literal bb 0))
  ([bb n]
   (let [more? (= 1 (.get bb))
         value (+ (* n 16) (read-number 4 bb))]
     (if more? (recur bb value) value))))

(declare byte-buffer->packets)
(defn read-operator
  [bb]
  (if (= 0 (.get bb))
    (let [length (read-number 15 bb)
          sub-bb (doto (.slice bb) (.limit length))]
      (.position bb (+ (.position bb) length))
      (byte-buffer->packets sub-bb))
    (let [count (read-number 11 bb)]
      (byte-buffer->packets bb count))))

(defn byte-buffer->packet
  [bb]
  (let [version (read-number 3 bb)
        type-id (read-number 3 bb)
        packet {:version version :type-id type-id}]
    (if (= 4 type-id)
      (assoc packet :value (read-literal bb))
      (assoc packet :sub-packets (read-operator bb)))))

(defn byte-buffer->packets
  ([bb]
   (if (.hasRemaining bb)
     (cons (byte-buffer->packet bb) (byte-buffer->packets bb))
     ()))
  ([bb n]
   (if (> n 0)
     (cons (byte-buffer->packet bb) (byte-buffer->packets bb (dec n)))
     ())))

(def string->packet
  (comp byte-buffer->packet string->byte-buffer))

(defn version-sum
  [{version :version sub-packets :sub-packets}]
  (apply + version (map version-sum (or sub-packets ()))))

;(println (string->packet "D2FE28"))
;(println (string->packet "38006F45291200"))
;(println (string->packet "EE00D40C823060"))
;(println (string->packet "8A004A801A8002F478"))
;(println (version-sum (string->packet "8A004A801A8002F478")))
;(println (string->packet "620080001611562C8802118E34"))
;(println (version-sum (string->packet "620080001611562C8802118E34")))
;(println (string->packet "C0015000016115A2E0802F182340"))
;(println (version-sum (string->packet "C0015000016115A2E0802F182340")))
;(println (string->packet "A0016C880162017C3686B18A3D4780"))
;(println (version-sum (string->packet "A0016C880162017C3686B18A3D4780")))

;(println (string->packet input))
(println (version-sum (string->packet input)))
