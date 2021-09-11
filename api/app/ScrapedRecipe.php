<?php

namespace App;

use Illuminate\Support\Collection;

class ScrapedRecipe extends Recipe
{
    /**
     * @var array $data
     */
    protected $data;

    public function __construct(array $data)
    {
        $this->data = $data;
    }

    /**
     * @return string
     */
    protected function title(): string
    {
        return $this->data['title'] ?? '';
    }

    /**
     * @return int
     */
    protected function time(): int
    {
        return $this->data['time'] ?? 0;
    }

    /**
     * @return string
     */
    protected function servings(): string
    {
        return $this->data['servings'] ?? '';
    }

    /**
     * @return Collection
     */
    protected function ingredients(): Collection
    {
        return collect($this->data['ingredients']) ?? collect();
    }

    /**
     * @return Collection
     */
    protected function instructions(): Collection
    {
        return collect($this->data['instructions']) ?? collect();
    }

    /**
     * @return string
     */
    protected function image(): string
    {
        return $this->data['image'] ?? '';
    }

    /**
     * @return string
     */
    protected function url(): string
    {
        return $this->data['url'] ?? '';
    }
}
