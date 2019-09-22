package com.example.demo.Controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;

@Controller
public class DemoController {

    @GetMapping("/")
    public String index(
            @RequestParam(name="file", required = false) String file,
            Model model) {

        model.addAttribute("file", file);

        return "index";
    }
}
