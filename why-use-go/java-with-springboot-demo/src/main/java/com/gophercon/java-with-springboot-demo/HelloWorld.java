package com.gophercon.springbootdemo;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Arrays;


@RestController
public class HelloWorld {

    @RequestMapping("/")
    public String index(String input) {
        return "Spring received: " + input;
    }

}
