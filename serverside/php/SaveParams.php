<?php

class SaveParam
{
    private $domain;

    private $clickParam;

    private $keyName;

    private $cookieExpire;

    public function __construct()
    {
    }

    public function saveCookieParams(): bool
    {
        setcookie($this->keyName, $this->clickParam, $this->cookieExpire,  '/', $this->domain, true, true);

        return true;
    }
}
