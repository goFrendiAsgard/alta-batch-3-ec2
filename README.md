# MVC

Model-View-Controller

# Interface

* To create `PersonController`, you need a `PersonModel`.
* `PersonModel` is an interface.
* `PersonModel` is anything that comply `PersonModel` interface (have bunch of functions defined in the interface)
* We have two `PersonModel`:
    - `PersonMemModel`: This one save the data in memory (as list of person)
    - `PersonDbModel`: This one use gorm