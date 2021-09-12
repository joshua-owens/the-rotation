<?php

namespace App;

use Illuminate\Support\Facades\Http;

class RecipeScraper implements Scraper
{
    /**
     * @var string
     */
    private $apiUrl;

    public function __construct()
    {
        $this->apiUrl = config_path('scraper.url');
    }

    public function scrape(string $url): Recipe
    {
        $endpoint = "$this->apiUrl/scrape";

        $data = Http::post($endpoint, [
            'url' => $url,
        ])
            ->json();

        return new ScrapedRecipe($data);
    }
}
