<?php

namespace App;

use Illuminate\Support\Collection;

/**
 * @property-read string $title
 * @property-read string $time
 * @property-read string $servings
 * @property-read Collection $ingredients
 * @property-read Collection instructions
 * @property-read string $image
 * @property-read string $url
 */
abstract class Recipe {
    abstract public function title(): string;
    abstract public function time(): int;
    abstract public function servings(): string;
    abstract public function ingredients(): Collection;
    abstract public function instructions(): Collection;
    abstract public function image(): string;
    abstract public function url(): string;

    public function __get(string $name)
    {
        return $this->{$name}();
    }
}
