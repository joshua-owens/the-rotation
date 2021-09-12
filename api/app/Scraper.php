<?php

namespace App;

interface Scraper {
    public function scrape(string $url): Recipe;
}
