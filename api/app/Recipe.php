<?php

namespace App;

abstract class Recipe {
    abstract public function title(): string;
    abstract public function time(): string;
    abstract public function servings(): string;
    abstract public function ingredients(): string;
    abstract public function instructions(): string;
    abstract public function image(): string;
    abstract public function url(): string;

    public function __get(string $name)
    {
        return $this->{$name}();
    }
}
