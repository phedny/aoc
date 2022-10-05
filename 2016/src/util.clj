(ns util (:require [clojure.string :as str]))

(defn input-basename
  []
  (if-let [v (ns-resolve *ns* (symbol "filename"))]
    (var-get v)
    "real"))

(defn input-filename
  []
  (let [day-number (get (re-find (re-matcher #"day-(\d+)" (name (ns-name *ns*)))) 1)]
    (str "../inputs/" day-number "/" (input-basename) ".txt")))

(defn apply-fns
  [fs xs]
  (map #(%1 %2) fs xs))

(defn read-input
  ([] (read-input #"\n" identity))
  ([fn-or-re] (cond
                (fn? fn-or-re) (read-input #"\n" fn-or-re)
                (instance? java.util.regex.Pattern fn-or-re) (read-input fn-or-re identity)))
  ([line-re f] (map f (str/split (slurp (input-filename)) line-re)))
  ([line-re column-re f & fs] (read-input line-re #(apply-fns (apply vector f fs) (str/split % column-re)))))
