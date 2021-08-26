<?php

/**
 * Measure using cookies
 */
class MeasureCookie
{
    /** @var string Domain name */
    private $domain;

    /** @var string Click parameters attached to the URL */
    private $clickParam;

    /** @var string Save key name */
    private $keyName;

    /** @var string cookie expire time */
    private $cookieExpire;

    /**
     * Set setcookie() params;
     * @param string $domain
     * @param string $clickParam
     * @param string $keyName
     * @param string $cookieExpire
     */
    public function __construct(string $domain, string $clickParam, string $keyName, string $cookieExpire)
    {
        $this->domain = $domain;
        $this->clickParam = $clickParam;
        $this->keyName = $keyName;
        $this->cookieExpire = $cookieExpire;
    }

    /**
     * Set param value to cookie 
     * @return boolean
     */
    public function saveCookieParams(): bool
    {
        // secure on, OnlyHttp on 
        return setcookie($this->keyName, $this->clickParam, $this->cookieExpire,  '/', $this->domain, true, true);
    }

    /**
     * Send param value to measurement URL
     * @return boolean
     */
    public function sendCookieParams(): bool
    {
        $cookieParam = _COOKIE[$this->keyName];

        // this is sample measurement URL.
        $url = 'https://apiurl.com?param=' . $cookieParam;

        // @TODO Use_guzzle and guzzle Exception
        $ch = curl_init();

        curl_setopt_array(
            $ch,
            [
                CURLOPT_URL => $url,
                CURLOPT_RETURNTRANSFER => true,
            ]
        );

        try {
            $curlResult = curl_exec($ch);
            curl_close($ch);

            // @TODO Use GuzzleException
            if ($curlResult === false) {
                throw new Exception($curlResult);
            }
        } catch (\Exception $e) {
            // Exception handling..
            return false;
        }

        return true;
    }
}
