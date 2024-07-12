package handlers

import (
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/jwt"
    "time"
)

type Config struct {
    Log struct {
        Level     string `json:"level"`
        Timestamp bool   `json:"timestamp"`
    } `json:"log"`
    DNS struct {
        Servers []struct {
            Tag     string `json:"tag"`
            Address string `json:"address"`
            Detour  string `json:"detour,omitempty"`
        } `json:"servers"`
        Rules []struct {
            Geosite     []string `json:"geosite,omitempty"`
            Server      string   `json:"server"`
            DisableCache bool    `json:"disable_cache,omitempty"`
            Outbound    []string `json:"outbound,omitempty"`
        } `json:"rules"`
    } `json:"dns"`
}

func GetConfig(ctx iris.Context) {
    user := ctx.Values().Get("jwt").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    userID := int(claims["user_id"].(float64))

    // Here we should check if the user ID is valid and exists, for simplicity assuming it's valid
    expirationTime := time.Now().Add(10 * time.Minute).Unix()

    config := Config{
        Log: struct {
            Level     string `json:"level"`
            Timestamp bool   `json:"timestamp"`
        }{
            Level:     "debug",
            Timestamp: true,
        },
        DNS: struct {
            Servers []struct {
                Tag     string `json:"tag"`
                Address string `json:"address"`
                Detour  string `json:"detour,omitempty"`
            } `json:"servers"`
            Rules []struct {
                Geosite     []string `json:"geosite,omitempty"`
                Server      string   `json:"server"`
                DisableCache bool    `json:"disable_cache,omitempty"`
                Outbound    []string `json:"outbound,omitempty"`
            } `json:"rules"`
        }{
            Servers: []struct {
                Tag     string `json:"tag"`
                Address string `json:"address"`
                Detour  string `json:"detour,omitempty"`
            }{
                {Tag: "dns_direct", Address: "local", Detour: "direct"},
                {Tag: "dns_proxy", Address: "tls://1.1.1.1", Detour: "direct"},
                {Tag: "dns_block", Address: "rcode://success"},
            },
            Rules: []struct {
                Geosite     []string `json:"geosite,omitempty"`
                Server      string   `json:"server"`
                DisableCache bool    `json:"disable_cache,omitempty"`
                Outbound    []string `json:"outbound,omitempty"`
            }{
                {Geosite: []string{"category-ads-all"}, Server: "dns_block", DisableCache: true},
                {Geosite: []string{"category-games@cn"}, Server: "dns_direct"},
                {Outbound: []string{"any"}, Server: "dns_proxy"},
            },
        },
    }

    configWithExpiration := struct {
        Config
        Expiration int64 `json:"expiration"`
    }{
        Config:     config,
        Expiration: expirationTime,
    }

    ctx.JSON(configWithExpiration)
}

