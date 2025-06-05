package com.example.repository;

import com.example.entity.MetaAttr;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface MetaAttrRepository extends JpaRepository<MetaAttr, Long> {
    Page<MetaAttr> findByIsDeletedFalse(Pageable pageable);
} 