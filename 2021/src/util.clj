(ns util)

(defn input-basename
  []
  (if-let [v (ns-resolve *ns* (symbol "filename"))]
    (var-get v)
    "real"))

(defn input-filename
  []
  (str "../inputs/" (clojure.string/replace (ns-name *ns*) #"day-" "") "/" (input-basename) ".txt"))

(defn read-input
  ([] (clojure.string/split (slurp (input-filename)) #"\n"))
  ([f] (map f (read-input))))
