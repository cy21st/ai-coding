package com.example.entity;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Data;
import org.hibernate.annotations.SQLDelete;
import org.hibernate.annotations.Where;

import javax.persistence.*;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.List;

@Data
@Entity
@Table(name = "meta_event")
@SQLDelete(sql = "UPDATE meta_event SET is_deleted = true WHERE id = ?")
@Where(clause = "is_deleted = false")
public class MetaEvent {
    private static final DateTimeFormatter DATE_FORMATTER = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss");

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "event_name", length = 200, nullable = false)
    private String eventName;

    @Column(name = "event_desc", columnDefinition = "TEXT")
    private String eventDesc;

    @Column(name = "is_deleted", nullable = false)
    private Boolean isDeleted = false;

    @JsonIgnore
    @Column(name = "created_at", nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @JsonIgnore
    @Column(name = "updated_at", nullable = false)
    private LocalDateTime updatedAt;

    @ManyToMany(fetch = FetchType.LAZY)
    @JoinTable(
        name = "meta_relation",
        joinColumns = @JoinColumn(name = "event_id"),
        inverseJoinColumns = @JoinColumn(name = "attr_id")
    )
    @JsonIgnoreProperties("events")
    private List<MetaAttr> attributes;

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