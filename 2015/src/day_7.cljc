(ns day-7 (:require util [clojure.string :refer [split]]))

(def input
  (util/read-input #(split % #" -> ")))

(defn literal
  [l]
  (if (re-matches #"\d+" l)
    (Short/parseShort l)
    l))

(defn parse-expression
  [expression]
  (let [[_ a op b] (re-find (re-matcher #"(?:([a-z0-9]+) )?(?:(AND|OR|NOT|LSHIFT|RSHIFT) )?([a-z0-9]+)" expression))]
    (case op
      nil [identity (literal b)]
      "AND" [bit-and (literal a) (literal b)]
      "OR" [bit-or (literal a) (literal b)]
      "NOT" [bit-not (literal b)]
      "LSHIFT" [bit-shift-left (literal a) (literal b)]
      "RSHIFT" [unsigned-bit-shift-right (literal a) (literal b)])))

(def defined-vars
  (reduce #(assoc %1 (second %2) (parse-expression (first %2))) {} input))

(defn resolve-var
  [vars args]
  (cond
    (number? args) args
    (every? number? (rest args)) (apply (first args) (rest args))
    :else (apply vector (first args) (map #(if (number? (vars %)) (vars %) %) (rest args)))))

(defn resolve-vars
  [vars]
  (if (every? (comp number? second) vars)
    vars
    (recur (reduce #(assoc %1 (first %2) (resolve-var vars (second %2))) {} vars))))

(def resolved-vars
  (resolve-vars defined-vars))

(println (resolved-vars "a"))

(def defined-vars-2
  (assoc defined-vars "b" (resolved-vars "a")))

(def resolved-vars-2
  (resolve-vars defined-vars-2))

(println (resolved-vars-2 "a"))
