package com.example.config;

import com.example.entity.AdminUser;
import com.example.repository.AdminUserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.crypto.password.PasswordEncoder;

@Configuration
public class InitialDataConfig {

    @Autowired
    private AdminUserRepository adminUserRepository;

    @Autowired
    private PasswordEncoder passwordEncoder;

    @Bean
    public CommandLineRunner initializeData() {
        return args -> {
            // 检查是否已存在管理员用户
            if (!adminUserRepository.existsByUsername("admin")) {
                AdminUser adminUser = new AdminUser();
                adminUser.setUsername("admin");
                adminUser.setPassword(passwordEncoder.encode("123456"));
                adminUser.setRole("admin");
                adminUserRepository.save(adminUser);
                System.out.println("初始化管理员用户成功");
            }
        };
    }
} 