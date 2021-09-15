<?php

namespace App;

use App\Exceptions\FailedToScrapeUrlException;
use Illuminate\Support\Facades\Http;

class RecipeScraper implements Scraper
{
    /**
     * @var string
     */
    private $apiUrl;

    public function __construct()
    {
        $this->apiUrl = config('scraper.url');
    }

    /**
     * @param string $url
     * @return Recipe
     * @throws FailedToScrapeUrlException
     */
    public function scrape(string $url): Recipe
    {
        $endpoint = "$this->apiUrl/scrape/";

        $response = Http::asForm()->post($endpoint, [
            'url' => $url,
        ]);

        if ($response->failed()) {
            throw new FailedToScrapeUrlException;
        }

        return new ScrapedRecipe($response->json());
    }
}
