(ns util (:require [clojure.string :as str]))

(defn input-basename
  []
  (if-let [v (ns-resolve *ns* (symbol "filename"))]
    (var-get v)
    "real"))

(defn input-filename
  []
  (str "../inputs/" (str/replace (ns-name *ns*) #"day-" "") "/" (input-basename) ".txt"))

(defn apply-fns
  [fs xs]
  (map #(%1 %2) fs xs))

(defn read-input
  ([] (str/split (slurp (input-filename)) #"\n"))
  ([f] (map f (read-input)))
  ([f & fs] (read-input #(apply-fns (apply vector f fs) (str/split % #" ")))))
