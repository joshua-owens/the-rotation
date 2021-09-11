<?php

namespace App;

use Illuminate\Support\Collection;

class ScrapedRecipe extends Recipe
{
    /**
     * @var array $data
     */
    private $data;

    public function __construct(array $data)
    {
        $this->data = $data;
    }

    /**
     * @return string
     */
    public function title(): string
    {
        return $this->data['title'] ?? '';
    }

    /**
     * @return int
     */
    public function time(): int
    {
        return $this->data['time'] ?? 0;
    }

    /**
     * @return string
     */
    public function servings(): string
    {
        return $this->data['servings'] ?? '';
    }

    /**
     * @return Collection
     */
    public function ingredients(): Collection
    {
        return collect($this->data['ingredients']) ?? collect();
    }

    /**
     * @return Collection
     */
    public function instructions(): Collection
    {
        return collect($this->data['instructions']) ?? collect();
    }

    /**
     * @return string
     */
    public function image(): string
    {
        return $this->data['image'] ?? '';
    }

    /**
     * @return string
     */
    public function url(): string
    {
        return $this->data['url'] ?? '';
    }
}
