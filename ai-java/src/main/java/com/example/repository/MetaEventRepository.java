package com.example.repository;

import com.example.entity.MetaEvent;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;

public interface MetaEventRepository extends JpaRepository<MetaEvent, Long> {
    Page<MetaEvent> findByIsDeletedFalse(Pageable pageable);
} 