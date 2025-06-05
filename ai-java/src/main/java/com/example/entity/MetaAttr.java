package com.example.entity;

import com.fasterxml.jackson.annotation.JsonIgnore;
import lombok.Data;
import org.hibernate.annotations.SQLDelete;
import org.hibernate.annotations.Where;

import javax.persistence.*;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.List;

@Data
@Entity
@Table(name = "meta_attr")
@SQLDelete(sql = "UPDATE meta_attr SET is_deleted = true WHERE id = ?")
@Where(clause = "is_deleted = false")
public class MetaAttr {
    private static final DateTimeFormatter DATE_FORMATTER = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss");

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "attr_name", length = 200, nullable = false)
    private String attrName;

    @Column(name = "attr_type", length = 50, nullable = false)
    private String attrType = "string";

    @Column(name = "attr_desc", columnDefinition = "TEXT")
    private String attrDesc;

    @Column(name = "is_deleted", nullable = false)
    private Boolean isDeleted = false;

    @JsonIgnore
    @Column(name = "created_at", nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @JsonIgnore
    @Column(name = "updated_at", nullable = false)
    private LocalDateTime updatedAt;

    @JsonIgnore
    @ManyToMany(mappedBy = "attributes")
    private List<MetaEvent> events;

    @Transient
    private String createdStr;

    @Transient
    private String updatedStr;

    @PrePersist
    protected void onCreate() {
        createdAt = LocalDateTime.now();
        updatedAt = LocalDateTime.now();
    }

    @PreUpdate
    protected void onUpdate() {
        updatedAt = LocalDateTime.now();
    }

    @PostLoad
    protected void onLoad() {
        if (createdAt != null) {
            createdStr = createdAt.format(DATE_FORMATTER);
        }
        if (updatedAt != null) {
            updatedStr = updatedAt.format(DATE_FORMATTER);
        }
    }
} 