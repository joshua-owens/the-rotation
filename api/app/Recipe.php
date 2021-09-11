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
    abstract protected function title(): string;
    abstract protected function time(): int;
    abstract protected function servings(): string;
    abstract protected function ingredients(): Collection;
    abstract protected function instructions(): Collection;
    abstract protected function image(): string;
    abstract protected function url(): string;

    public function __get(string $name)
    {
        return $this->{$name}();
    }
}
